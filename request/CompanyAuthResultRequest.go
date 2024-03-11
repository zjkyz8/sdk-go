package request

import (
	"qiyuesuo/sdk/http"
)

type CompanyAuthResultRequest struct {
	// 认证请求ID
	RequestId string `json:"requestId,omitempty"`
	// 待查询的认证公司名称
	CompanyName string `json:"companyName,omitempty"`
}

func (obj CompanyAuthResultRequest) GetUrl() string {
	return "/companyauth/result"
}

func (obj CompanyAuthResultRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewGetHttpParameter()
	parameter.AddParam("requestId", obj.RequestId)
	parameter.AddParam("companyName", obj.CompanyName)
	return parameter
}
