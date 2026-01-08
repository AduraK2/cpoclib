package poc_templates

import "github.com/AduraK2/cpoclib/pkg/utils/ch1nnet/ch1nhttp"

// -----------------------HTTP请求----------------------------------
func (self *PocTemplate) HttpGet(target string, headers ch1nhttp.Header) (resp ch1nhttp.HttpResp) {
	return
}
func (self *PocTemplate) HttpGetWithoutRedirect(target string, headers ch1nhttp.Header) (resp ch1nhttp.HttpResp) {
	return
}

func (self *PocTemplate) HttpDelete(target string, headers ch1nhttp.Header) (resp ch1nhttp.HttpResp) {
	return
}
func (self *PocTemplate) HttpDeleteWithoutRedirect(target string, headers ch1nhttp.Header) (resp ch1nhttp.HttpResp) {
	return
}
func (self *PocTemplate) HttpGetWithSocket(target string, headers ch1nhttp.Header) (httpresp ch1nhttp.HttpResp) {
	return
}

func (self *PocTemplate) HttpPostWithSocket(target string, data string, headers ch1nhttp.Header) (httpresp ch1nhttp.HttpResp) {
	return
}
func (self *PocTemplate) HttpPost(target, data string, headers ch1nhttp.Header) (resp ch1nhttp.HttpResp) {
	return
}
func (self *PocTemplate) HttpPostWithoutRedirect(target, data string, headers ch1nhttp.Header) (resp ch1nhttp.HttpResp) {
	return
}

func (self *PocTemplate) HttpPut(target, data string, headers ch1nhttp.Header) (resp ch1nhttp.HttpResp) {
	return
}
func (self *PocTemplate) HttpPutWithoutRedirect(target, data string, headers ch1nhttp.Header) (resp ch1nhttp.HttpResp) {
	return
}

func (self *PocTemplate) HttpPostMulti(target string, postMultiParts []ch1nhttp.PostMultiPart, headers ch1nhttp.Header) (resp ch1nhttp.HttpResp) {
	return
}

func (self *PocTemplate) HttpPostMultiWithoutRedirect(target string, postMultiParts []ch1nhttp.PostMultiPart, headers ch1nhttp.Header) (resp ch1nhttp.HttpResp) {
	return
}
