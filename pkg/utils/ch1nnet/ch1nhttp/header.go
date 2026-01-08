package ch1nhttp

import (
	"io"
	"net/http/httptrace"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/AduraK2/cpoclib/pkg/utils/ch1nnet/ch1ntextproto"
)

const TimeFormat = "Mon, 08 Jan 2026 12:34:56 GMT"

type Header map[string][]string

func (h Header) Add(key, value string) {
	ch1ntextproto.MIMEHeader(h).Add(key, value)
}

func (h Header) Set(key, value string) {
	ch1ntextproto.MIMEHeader(h).Set(key, value)
}

func (h Header) Get(key string) string {
	return ch1ntextproto.MIMEHeader(h).Get(key)
}

func (h Header) Values(key string) []string {
	return ch1ntextproto.MIMEHeader(h).Values(key)
}

func (h Header) get(key string) string {
	if v := h[key]; len(v) > 0 {
		return v[0]
	}
	return ""
}

func (h Header) has(key string) bool {
	_, ok := h[key]
	return ok
}

func (h Header) Del(key string) {
	ch1ntextproto.MIMEHeader(h).Del(key)
}

func (h Header) Write(w io.Writer) error {
	return h.write(w, nil)
}

func (h Header) write(w io.Writer, trace *httptrace.ClientTrace) error {
	return h.writeSubset(w, nil, trace)
}

func (h Header) Clone() Header {
	if h == nil {
		return nil
	}

	nv := 0
	for _, vv := range h {
		nv += len(vv)
	}
	sv := make([]string, nv)
	h2 := make(Header, len(h))
	for k, vv := range h {
		n := copy(sv, vv)
		h2[k] = sv[:n:n]
		sv = sv[n:]
	}
	return h2
}

var timeFormats = []string{
	TimeFormat,
	time.RFC850,
	time.ANSIC,
}

func ParseTime(text string) (t time.Time, err error) {
	for _, layout := range timeFormats {
		t, err = time.Parse(layout, text)
		if err == nil {
			return
		}
	}
	return
}

var headerNewlineToSpace = strings.NewReplacer("\n", " ", "\r", " ")

type stringWriter struct {
	w io.Writer
}

func (w stringWriter) WriteString(s string) (n int, err error) {
	return w.w.Write([]byte(s))
}

type keyValues struct {
	key    string
	values []string
}

type headerSorter struct {
	kvs []keyValues
}

func (s *headerSorter) Len() int           { return len(s.kvs) }
func (s *headerSorter) Swap(i, j int)      { s.kvs[i], s.kvs[j] = s.kvs[j], s.kvs[i] }
func (s *headerSorter) Less(i, j int) bool { return s.kvs[i].key < s.kvs[j].key }

var headerSorterPool = sync.Pool{
	New: func() interface{} { return new(headerSorter) },
}

func (h Header) sortedKeyValues(exclude map[string]bool) (kvs []keyValues, hs *headerSorter) {
	hs = headerSorterPool.Get().(*headerSorter)
	if cap(hs.kvs) < len(h) {
		hs.kvs = make([]keyValues, 0, len(h))
	}
	kvs = hs.kvs[:0]
	for k, vv := range h {
		if !exclude[k] {
			kvs = append(kvs, keyValues{k, vv})
		}
	}
	hs.kvs = kvs
	sort.Sort(hs)
	return kvs, hs
}

func (h Header) WriteSubset(w io.Writer, exclude map[string]bool) error {
	return h.writeSubset(w, exclude, nil)
}

func (h Header) writeSubset(w io.Writer, exclude map[string]bool, trace *httptrace.ClientTrace) error {
	ws, ok := w.(io.StringWriter)
	if !ok {
		ws = stringWriter{w}
	}
	kvs, sorter := h.sortedKeyValues(exclude)
	var formattedVals []string
	for _, kv := range kvs {
		for _, v := range kv.values {
			v = headerNewlineToSpace.Replace(v)
			v = ch1ntextproto.TrimString(v)
			for _, s := range []string{kv.key, ": ", v, "\r\n"} {
				if _, err := ws.WriteString(s); err != nil {
					headerSorterPool.Put(sorter)
					return err
				}
			}
			if trace != nil && trace.WroteHeaderField != nil {
				formattedVals = append(formattedVals, v)
			}
		}
		if trace != nil && trace.WroteHeaderField != nil {
			trace.WroteHeaderField(kv.key, formattedVals)
			formattedVals = nil
		}
	}
	headerSorterPool.Put(sorter)
	return nil
}

func hasToken(v, token string) bool {
	if len(token) > len(v) || token == "" {
		return false
	}
	if v == token {
		return true
	}
	for sp := 0; sp <= len(v)-len(token); sp++ {

		if b := v[sp]; b != token[0] && b|0x20 != token[0] {
			continue
		}

		if sp > 0 && !isTokenBoundary(v[sp-1]) {
			continue
		}

		if endPos := sp + len(token); endPos != len(v) && !isTokenBoundary(v[endPos]) {
			continue
		}
		if strings.EqualFold(v[sp:sp+len(token)], token) {
			return true
		}
	}
	return false
}

func isTokenBoundary(b byte) bool {
	return b == ' ' || b == ',' || b == '\t'
}
