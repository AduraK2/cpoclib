package exp_templates

import (
	"github.com/AduraK2/cpoclib/pkg/templates/exp_model"
	"github.com/AduraK2/cpoclib/pkg/utils/ch1nnet/ch1nhttp"
)

type ExpTemplate struct {
	Params exp_model.ExpSendParams
}

// 获取设置的http头部
func (self *ExpTemplate) GetInitExpHeaders() (headers ch1nhttp.Header) {
	return
}
