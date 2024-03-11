package request

import (
	"qiyuesuo/sdk/http"
)

type SubCompanyRemoveRequest struct {
	// 将要移除的子公司名称
	CompanyName string `json:"companyName,omitempty"`
}

func (obj SubCompanyRemoveRequest) GetUrl() string {
	return "/v2/subcompany/remove"
}

func (obj SubCompanyRemoveRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewGetHttpParameter()
	parameter.AddParam("companyName", obj.CompanyName)
	return parameter
}
