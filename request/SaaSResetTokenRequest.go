package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
)

type SaaSResetTokenRequest struct {
	CompanyId   string `json:"companyId,omitempty"`
	AccessToken string `json:"accessToken,omitempty"`
}

func (obj SaaSResetTokenRequest) GetUrl() string {
	return "/saas/v2/company/resettoken"
}

func (obj SaaSResetTokenRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
