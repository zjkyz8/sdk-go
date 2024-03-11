package request

import (
	"qiyuesuo/sdk/http"
)

type ContractShorturlRequest struct {
	ContractId string `json:"contractId,omitempty"`
	BizId      string `json:"bizId,omitempty"`
	TenantName string `json:"tenantName,omitempty"`
	Contact    string `json:"contact,omitempty"`
}

func (obj ContractShorturlRequest) GetUrl() string {
	return "/v2/contract/shorturl"
}

func (obj ContractShorturlRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewGetHttpParameter()
	parameter.AddParam("tenantName", obj.TenantName)
	parameter.AddParam("contractId", obj.ContractId)
	parameter.AddParam("contact", obj.Contact)
	parameter.AddParam("contractId", obj.ContractId)
	return parameter
}
