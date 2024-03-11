package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
	"qiyuesuo/sdk/model"
)

type SaaSPrivilegeMiniappTicketRequest struct {
	AppId            string      `json:"appId,omitempty"`
	CompanyId        string      `json:"companyId,omitempty"`
	User             *model.User `json:"user,omitempty"`
	CreateToken      *bool       `json:"createToken,omitempty"`
	SuccessUrl       string      `json:"successUrl,omitempty"`
	CallbackUrl      string      `json:"callbackUrl,omitempty"`
	PrivilegeModules []string    `json:"privilegeModules,omitempty"`
}

func (obj SaaSPrivilegeMiniappTicketRequest) GetUrl() string {
	return "/saas/v2/privilege/miniappexchange"
}

func (obj SaaSPrivilegeMiniappTicketRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
