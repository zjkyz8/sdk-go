package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
)

type TemplateStatusManageRequest struct {
	TemplateId string `json:"templateId,omitempty"`
	// ENABLE（“启用”）、DISABLE（“停用”）
	Operate string `json:"operate,omitempty"`
}

func (obj TemplateStatusManageRequest) GetUrl() string {
	return "/v2/template/status"
}

func (obj TemplateStatusManageRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
