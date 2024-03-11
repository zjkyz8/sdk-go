package request

import (
	"qiyuesuo/sdk/http"
)

type SaaSCompanyDetailRequest struct {
	CompanyName string `json:"companyName,omitempty"`
	RegisterNo  string `json:"registerNo,omitempty"`
}

func (obj SaaSCompanyDetailRequest) GetUrl() string {
	return "/saas/v2/company/detail"
}

func (obj SaaSCompanyDetailRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewGetHttpParameter()
	parameter.AddParam("companyName", obj.CompanyName)
	parameter.AddParam("registerNo", obj.RegisterNo)
	return parameter
}
