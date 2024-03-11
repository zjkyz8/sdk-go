package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
	"qiyuesuo/sdk/model"
)

type SaasPrivilegeUrlRequest struct {
	AppId            string      `json:"appId,omitempty"`
	CompanyId        string      `json:"companyId,omitempty"`
	User             *model.User `json:"user,omitempty"`
	CreateToken      *bool       `json:"createToken,omitempty"`
	SuccessUrl       string      `json:"successUrl,omitempty"`
	CallbackUrl      string      `json:"callbackUrl,omitempty"`
	PrivilegeModules []string    `json:"privilegeModules,omitempty"`
	CanShowSecret    *bool       `json:"canShowSecret,omitempty"`
	Lang             string      `json:"lang,omitempty"`
	PageStyleInfo    string      `json:"pageStyleInfo,omitempty"`
}

func (obj SaasPrivilegeUrlRequest) GetUrl() string {
	return "/saas/v2/privilege/url"
}

func (obj SaasPrivilegeUrlRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
