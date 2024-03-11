package request

import (
	"qiyuesuo/sdk/http"
)

type SaaSCompanyAuthResultRequest struct {
	CompanyName string `json:"companyName,omitempty"`
	RequestId   string `json:"requestId,omitempty"`
}

func (obj SaaSCompanyAuthResultRequest) GetUrl() string {
	return "/saas/v2/companyauth/result"
}

func (obj SaaSCompanyAuthResultRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewGetHttpParameter()
	parameter.AddParam("companyName", obj.CompanyName)
	parameter.AddParam("requestId", obj.RequestId)
	return parameter
}
