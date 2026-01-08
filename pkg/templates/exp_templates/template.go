package exp_templates

import (
	"github.com/AduraK2/cpoclib/pkg/templates/exp_model"
	"github.com/AduraK2/cpoclib/pkg/utils/ch1nnet/ch1nhttp"
)

type ExpTemplate struct {
	Params exp_model.ExpSendParamS
	//Num2AliasMap map[string] string
}

// 获取设置的http头部
func (self *ExpTemplate) GetInitExpHeaders() (headers ch1nhttp.Header) {
	return
}

// 安全获取map里的value
func (self *ExpTemplate) GetItemSafe(s map[string]string, key string) (ret string) {
	return
}

// 只有当前URL没有路径/目录时，会添加URI
func (self *ExpTemplate) AddUri(target, uri string) (result string) {

	return
}

// 追加URL，基于当前目录
func (self *ExpTemplate) AppendUri(target, uri string) (result string) {
	return
}
func (self *ExpTemplate) GetHostname(target string) (hostname string) {
	return
}
