package request

import (
	"qiyuesuo/sdk/http"
)

type TemplateRemoveRequest struct {
	// 模板ID
	TemplateId string `json:"templateId,omitempty"`
}

func (obj TemplateRemoveRequest) GetUrl() string {
	return "/v2/template/remove"
}

func (obj TemplateRemoveRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewGetHttpParameter()
	parameter.AddParam("templateId", obj.TemplateId)
	return parameter
}
