package exp_templates

import (
	"github.com/AduraK2/cpoclib/pkg/utils/ch1nnet/ch1nhttp"
)

// -----------------------HTTP请求----------------------------------
func (self *ExpTemplate) HttpGet(target string, headers ch1nhttp.Header) (resp ch1nhttp.HttpResp) {
	return
}
func (self *ExpTemplate) HttpGetWithoutRedirect(target string, headers ch1nhttp.Header) (resp ch1nhttp.HttpResp) {
	return
}

func (self *ExpTemplate) HttpDelete(target string, headers ch1nhttp.Header) (resp ch1nhttp.HttpResp) {
	return
}
func (self *ExpTemplate) HttpDeleteWithoutRedirect(target string, headers ch1nhttp.Header) (resp ch1nhttp.HttpResp) {
	return
}
func (self *ExpTemplate) HttpGetWithSocket(target string, headers ch1nhttp.Header) (httpresp ch1nhttp.HttpResp) {
	return
}

func (self *ExpTemplate) HttpPostWithSocket(target string, data string, headers ch1nhttp.Header) (httpresp ch1nhttp.HttpResp) {
	return
}
func (self *ExpTemplate) HttpPost(target, data string, headers ch1nhttp.Header) (resp ch1nhttp.HttpResp) {
	return
}
func (self *ExpTemplate) HttpPostWithoutRedirect(target, data string, headers ch1nhttp.Header) (resp ch1nhttp.HttpResp) {
	return
}

func (self *ExpTemplate) HttpPut(target, data string, headers ch1nhttp.Header) (resp ch1nhttp.HttpResp) {
	return
}
func (self *ExpTemplate) HttpPutWithoutRedirect(target, data string, headers ch1nhttp.Header) (resp ch1nhttp.HttpResp) {
	return
}

func (self *ExpTemplate) HttpPostMulti(target string, postMultiParts []ch1nhttp.PostMultiPart, headers ch1nhttp.Header) (resp ch1nhttp.HttpResp) {
	return
}
func (self *ExpTemplate) HttpPostMultiWithoutRedirect(target string, postMultiParts []ch1nhttp.PostMultiPart, headers ch1nhttp.Header) (resp ch1nhttp.HttpResp) {
	return
}
