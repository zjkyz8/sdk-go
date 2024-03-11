package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
	"qiyuesuo/sdk/model"
)

type ContractInvalidRequest struct {
	// 合同ID与业务ID二选一，不能同时为空
	ContractId string `json:"contractId,omitempty"`
	// 合同ID与业务ID二选一，不能同时为空
	BizId string `json:"bizId,omitempty"`
	// 若使用业务ID撤回作废合同，且合同是以子公司身份创建的，则需要传递该值，用于确定合同主体
	TenantName string `json:"tenantName,omitempty"`
	// 发起方签署作废文件时指定的印章，作废合同时使用。若发起方已签署，不传默认取发起方印章ID
	SealId string `json:"sealId,omitempty"`
	Reason string `json:"reason,omitempty"`
	// 默认false
	DeleteDoc *bool `json:"deleteDoc,omitempty"`
	// 默认true
	AutoSign *bool `json:"autoSign,omitempty"`
	// 仅单方作废时生效；传入的参数里联系方式重复，以第一条为准
	Receivers []*model.CopySendReceiver `json:"receivers,omitempty"`
}

func (obj ContractInvalidRequest) GetUrl() string {
	return "/v2/contract/invalid"
}

func (obj ContractInvalidRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
