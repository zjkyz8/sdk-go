package request

import (
	"qiyuesuo/sdk/http"
)

type CompanyDetailRequest struct {
	// 公司名称,registerNo为空时必传
	CompanyName string `json:"companyName,omitempty"`
	// 企业社会信用代码,companyName为空时必传
	RegisterNo string `json:"registerNo,omitempty"`
}

func (obj CompanyDetailRequest) GetUrl() string {
	return "/v2/company/detail"
}

func (obj CompanyDetailRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewGetHttpParameter()
	parameter.AddParam("companyName", obj.CompanyName)
	parameter.AddParam("registerNo", obj.RegisterNo)
	return parameter
}
