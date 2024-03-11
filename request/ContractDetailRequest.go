package request

import (
	"qiyuesuo/sdk/http"
)

type ContractDetailRequest struct {
	// 是否查询关联合同
	QueryRelatedContract string `json:"queryRelatedContract,omitempty"`
	// 子公司名称，若使用业务ID查询合同详情，且合同是以子公司身份创建的，则需要传递该值，用于确定合同主体
	TenantName string `json:"tenantName,omitempty"`
	// 合同ID，合同ID于业务ID二选一，不能同时为空
	ContractId string `json:"contractId,omitempty"`
	// 业务ID，合同ID于业务ID二选一，不能同时为空
	BizId string `json:"bizId,omitempty"`
	// 是否查询印章签署次数
	QuerySealStats string `json:"querySealStats,omitempty"`
	// 是否查询签署方签署位置信息
	QueryLocation string `json:"queryLocation,omitempty"`
	// 是否查询签署人的信息：如果签署动作已启动时返回待签署人，签署动作已完成时返回实际签署人；如果签署方是平台方或平台方的子公司，才会返回签署人信息
	QueryActionOperator string `json:"queryActionOperator,omitempty"`
}

func (obj ContractDetailRequest) GetUrl() string {
	return "/v2/contract/detail"
}

func (obj ContractDetailRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewGetHttpParameter()
	parameter.AddParam("queryRelatedContract", obj.QueryRelatedContract)
	parameter.AddParam("tenantName", obj.TenantName)
	parameter.AddParam("contractId", obj.ContractId)
	parameter.AddParam("bizId", obj.BizId)
	parameter.AddParam("querySealStats", obj.QuerySealStats)
	parameter.AddParam("queryLocation", obj.QueryLocation)
	parameter.AddParam("queryActionOperator", obj.QueryActionOperator)
	return parameter
}
