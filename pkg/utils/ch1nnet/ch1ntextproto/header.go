package ch1ntextproto

type MIMEHeader map[string][]string

func (h MIMEHeader) Add(key, value string) {
	h[key] = append(h[key], value)
}

func (h MIMEHeader) Set(key, value string) {
	h[key] = []string{value}
}

func (h MIMEHeader) Get(key string) string {
	if h == nil {
		return ""
	}
	v := h[key]
	if len(v) == 0 {
		return ""
	}
	return v[0]
}

func (h MIMEHeader) Values(key string) []string {
	if h == nil {
		return nil
	}
	return h[key]
}

func (h MIMEHeader) Del(key string) {
	delete(h, key)
}
