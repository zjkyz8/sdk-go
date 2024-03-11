package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
	"qiyuesuo/sdk/model"
)

type CustomerAddRequest struct {
	// COMPANY（企业）,PERSONAL（个人）
	TenantType string `json:"tenantType,omitempty"`
	// 当类型为COMPANY时必填
	CompanyName      string      `json:"companyName,omitempty"`
	User             *model.User `json:"user,omitempty"`
	AlternateContact string      `json:"alternateContact,omitempty"`
}

func (obj CustomerAddRequest) GetUrl() string {
	return "/v2/customer/add"
}

func (obj CustomerAddRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
