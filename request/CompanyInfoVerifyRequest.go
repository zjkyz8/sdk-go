package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
)

type CompanyInfoVerifyRequest struct {
	CompanyName string `json:"companyName,omitempty"`
	RegisterNo  string `json:"registerNo,omitempty"`
	LegalPerson string `json:"legalPerson,omitempty"`
}

func (obj CompanyInfoVerifyRequest) GetUrl() string {
	return "/v2/companyinfo/verify"
}

func (obj CompanyInfoVerifyRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
