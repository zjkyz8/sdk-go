package request

import (
	"qiyuesuo/sdk/http"
)

type SaaSCompanyAuthMiniappTicketRequest struct {
	CompanyName   string         `json:"companyName,omitempty"`
	ApplicantInfo string         `json:"applicantInfo,omitempty"`
	RegisterNo    string         `json:"registerNo,omitempty"`
	LegalPerson   string         `json:"legalPerson,omitempty"`
	License       *http.FileItem `json:"license,omitempty"`
	CallbackUrl   string         `json:"callbackUrl,omitempty"`
}

func (obj SaaSCompanyAuthMiniappTicketRequest) GetUrl() string {
	return "/saas/v2/companyauth/miniappexchange"
}

func (obj SaaSCompanyAuthMiniappTicketRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	parameter.AddParam("companyName", obj.CompanyName)
	parameter.AddParam("applicantInfo", obj.ApplicantInfo)
	parameter.AddParam("registerNo", obj.RegisterNo)
	parameter.AddParam("legalPerson", obj.LegalPerson)
	parameter.AddParam("callbackUrl", obj.CallbackUrl)
	parameter.AddFiles("license", obj.License)
	return parameter
}
