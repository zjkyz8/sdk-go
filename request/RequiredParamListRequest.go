package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
)

type RequiredParamListRequest struct {
	// 子公司名称，若使用业务ID查询合同详情，且合同是以子公司身份创建的，则需要传递该值，用于确定合同主体
	TenantName string `json:"tenantName,omitempty"`
	// 合同ID，合同ID于业务ID二选一，不能同时为空
	ContractId string `json:"contractId,omitempty"`
	// 业务ID，合同ID于业务ID二选一，不能同时为空
	BizId string `json:"bizId,omitempty"`
}

func (obj RequiredParamListRequest) GetUrl() string {
	return "/v2/contract/required/detail"
}

func (obj RequiredParamListRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
