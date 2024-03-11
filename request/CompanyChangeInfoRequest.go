package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
	"qiyuesuo/sdk/model"
)

type CompanyChangeInfoRequest struct {
	// 公司ID、公司名称、统一社会信用代码不能同时为空
	CompanyName string `json:"companyName,omitempty"`
	// 公司ID、公司名称、统一社会信用代码不能同时为空
	CompanyId *float64 `json:"companyId,omitempty"`
	// 公司ID、公司名称、统一社会信用代码不能同时为空
	RegisterNo string      `json:"registerNo,omitempty"`
	Applicant  *model.User `json:"applicant,omitempty"`
	// 默认为false(即申请PC端页面)
	H5 *bool `json:"h5,omitempty"`
}

func (obj CompanyChangeInfoRequest) GetUrl() string {
	return "/v2/company/changeinfo"
}

func (obj CompanyChangeInfoRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
