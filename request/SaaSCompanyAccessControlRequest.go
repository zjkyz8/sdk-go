package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
	"qiyuesuo/sdk/model"
)

type SaaSCompanyAccessControlRequest struct {
	Company *model.Company `json:"company,omitempty"`
	Operate string         `json:"operate,omitempty"`
}

func (obj SaaSCompanyAccessControlRequest) GetUrl() string {
	return "/saas/v2/company/accesscontrol"
}

func (obj SaaSCompanyAccessControlRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
