package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
	"qiyuesuo/sdk/model"
)

type ContractCopySendRequest struct {
	// 合同ID与业务ID二选一，不能同时为空
	ContractId string `json:"contractId,omitempty"`
	// 合同ID与业务ID二选一，不能同时为空
	BizId string `json:"bizId,omitempty"`
	// 若使用业务ID发起合同，且合同是以子公司身份创建的，则需要传递该值，用于确定合同主体
	TenantName string                    `json:"tenantName,omitempty"`
	Receivers  []*model.CopySendReceiver `json:"receivers,omitempty"`
}

func (obj ContractCopySendRequest) GetUrl() string {
	return "/v2/contract/copysend"
}

func (obj ContractCopySendRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
