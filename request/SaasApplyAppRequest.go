package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
)

type SaasApplyAppRequest struct {
	CompanyId                  string `json:"companyId,omitempty"`
	AppName                    string `json:"appName,omitempty"`
	Domain                     string `json:"domain,omitempty"`
	LoginUrl                   string `json:"loginUrl,omitempty"`
	CasVerifyUrl               string `json:"casVerifyUrl,omitempty"`
	CallbackUrl                string `json:"callbackUrl,omitempty"`
	ContractCallbackUrl        string `json:"contractCallbackUrl,omitempty"`
	MessageStrategy            string `json:"messageStrategy,omitempty"`
	ContractMessageCallbackUrl string `json:"contractMessageCallbackUrl,omitempty"`
	ContractByPage             *bool  `json:"contractByPage,omitempty"`
	CallbackForPage            *bool  `json:"callbackForPage,omitempty"`
	SealAuthCallbackUrl        string `json:"sealAuthCallbackUrl,omitempty"`
}

func (obj SaasApplyAppRequest) GetUrl() string {
	return "/saas/v2/company/applyapp"
}

func (obj SaasApplyAppRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
