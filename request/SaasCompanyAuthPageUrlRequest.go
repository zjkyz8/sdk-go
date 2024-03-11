package request

import (
	"qiyuesuo/sdk/http"
)

type SaasCompanyAuthPageUrlRequest struct {
	CompanyName   string         `json:"companyName,omitempty"`
	RegisterNo    string         `json:"registerNo,omitempty"`
	LegalPerson   string         `json:"legalPerson,omitempty"`
	ApplicantInfo string         `json:"applicantInfo,omitempty"`
	License       *http.FileItem `json:"license,omitempty"`
	CallbackUrl   string         `json:"callbackUrl,omitempty"`
	Lang          string         `json:"lang,omitempty"`
	PageStyleInfo string         `json:"pageStyleInfo,omitempty"`
}

func (obj SaasCompanyAuthPageUrlRequest) GetUrl() string {
	return "/saas/v2/companyauth/pageurl"
}

func (obj SaasCompanyAuthPageUrlRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	parameter.AddParam("companyName", obj.CompanyName)
	parameter.AddParam("registerNo", obj.RegisterNo)
	parameter.AddParam("legalPerson", obj.LegalPerson)
	parameter.AddParam("applicantInfo", obj.ApplicantInfo)
	parameter.AddFiles("license", obj.License)
	parameter.AddParam("callbackUrl", obj.CallbackUrl)
	parameter.AddParam("lang", obj.Lang)
	parameter.AddParam("pageStyleInfo", obj.PageStyleInfo)
	return parameter
}
