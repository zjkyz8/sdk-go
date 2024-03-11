package request

import (
	"qiyuesuo/sdk/http"
)

type TemplatePageRequest struct {
	// 模板ID
	TemplateId string `json:"templateId,omitempty"`
}

func (obj TemplatePageRequest) GetUrl() string {
	return "/v2/template/pageurl"
}

func (obj TemplatePageRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewGetHttpParameter()
	parameter.AddParam("templateId", obj.TemplateId)
	return parameter
}
