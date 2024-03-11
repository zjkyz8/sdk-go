package request

import (
	"qiyuesuo/sdk/http"
)

type TemplateDetailRequest struct {
	// 模板ID
	TemplateId string `json:"templateId,omitempty"`
}

func (obj TemplateDetailRequest) GetUrl() string {
	return "/v2/template/detail"
}

func (obj TemplateDetailRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewGetHttpParameter()
	parameter.AddParam("templateId", obj.TemplateId)
	return parameter
}
