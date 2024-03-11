package request

import (
	"qiyuesuo/sdk/http"
)

type ContractJoinUrlRequest struct {
	// 合同ID与业务ID二选一，不能同时为空
	ContractId string `json:"contractId,omitempty"`
	// 合同ID与业务ID二选一，不能同时为空
	BizId string `json:"bizId,omitempty"`
	// 若以子公司身份签署，需要传递子公司全名
	TenantName string `json:"tenantName,omitempty"`
	Type_      string `json:"type,omitempty"`
}

func (obj ContractJoinUrlRequest) GetUrl() string {
	return "/v2/contract/joinurl"
}

func (obj ContractJoinUrlRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewGetHttpParameter()
	parameter.AddParam("contractId", obj.ContractId)
	parameter.AddParam("bizId", obj.BizId)
	parameter.AddParam("tenantName", obj.TenantName)
	parameter.AddParam("type", obj.Type_)
	return parameter
}
