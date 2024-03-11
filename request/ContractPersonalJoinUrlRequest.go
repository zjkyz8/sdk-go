package request

import "qiyuesuo/sdk/http"

type ContractPersonalJoinUrlRequest struct {
	// 合同ID与业务ID二选一，不能同时为空
	ContractId string `json:"contractId,omitempty"`
	// 合同ID与业务ID二选一，不能同时为空
	BizId string `json:"bizId,omitempty"`
	// 若以子公司身份签署，需要传递子公司全名
	TenantName string `json:"tenantName,omitempty"`
}

func (obj ContractPersonalJoinUrlRequest) GetUrl() string {
	return "/v2/contract/personal/joinurl"
}

func (obj ContractPersonalJoinUrlRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewGetHttpParameter()
	parameter.AddParam("tenantName", obj.TenantName)
	parameter.AddParam("contractId", obj.ContractId)
	parameter.AddParam("bizId", obj.BizId)
	return parameter
}
