package request

import (
	"qiyuesuo/sdk/http"
)

type SaasPrivilegeDetailRequest struct {
	AppId     string `json:"appId,omitempty"`
	CompanyId string `json:"companyId,omitempty"`
}

func (obj SaasPrivilegeDetailRequest) GetUrl() string {
	return "/saas/v2/privilege/query"
}

func (obj SaasPrivilegeDetailRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewGetHttpParameter()
	parameter.AddParam("appId", obj.AppId)
	parameter.AddParam("companyId", obj.CompanyId)
	return parameter
}
