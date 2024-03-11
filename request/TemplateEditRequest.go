package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
	"qiyuesuo/sdk/model"
)

type TemplateEditRequest struct {
	TemplateId string `json:"templateId,omitempty"`
	// 参数为空时不作处理
	Title          string                 `json:"title,omitempty"`
	TemplateParams []*model.TemplateParam `json:"templateParams,omitempty"`
	// 替换此模板原有管理员，参数为空时不作处理
	Managers []*model.User `json:"managers,omitempty"`
	// 自定义设置具体的人，替换此模板原有使用人，参数为空时不作处理
	Range_ []*model.User `json:"range,omitempty"`
}

func (obj TemplateEditRequest) GetUrl() string {
	return "/v2/template/edit"
}

func (obj TemplateEditRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
