package request

import (
	"qiyuesuo/sdk/http"
)

type ContractStreamRequest struct {
	// 子公司名称，若使用业务ID查询，且合同是以子公司身份创建的，则需要传递该值，用于确定合同主体
	TenantName string `json:"tenantName,omitempty"`
	// 合同ID，合同ID于业务ID二选一，不能同时为空
	ContractId string `json:"contractId,omitempty"`
	// 务ID，合同ID于业务ID二选一，不能同时为空
	BizId string `json:"bizId,omitempty"`
	// 操作类型，支持传入字段:ALL(全部)，VIEW(查看）,MODIFY（修改），DOWNLOAD（下载），PRINT（打印）,默认ALL；多个字段用逗号隔开（例：VIEW,MODIFY）
	Operation string `json:"operation,omitempty"`
}

func (obj ContractStreamRequest) GetUrl() string {
	return "/v2/contract/stream"
}

func (obj ContractStreamRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewGetHttpParameter()
	parameter.AddParam("tenantName", obj.TenantName)
	parameter.AddParam("contractId", obj.ContractId)
	parameter.AddParam("bizId", obj.BizId)
	parameter.AddParam("operation", obj.Operation)
	return parameter
}
