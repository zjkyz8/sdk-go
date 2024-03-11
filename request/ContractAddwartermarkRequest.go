package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
	"qiyuesuo/sdk/model"
)

type ContractAddwartermarkRequest struct {
	// 合同ID与业务ID二选一，不能同时为空
	ContractId string `json:"contractId,omitempty"`
	// 合同ID与业务ID二选一，不能同时为空
	BizId string `json:"bizId,omitempty"`
	// 若使用业务ID标识合同，且合同是以子公司身份创建的，则需要传递该值，用于确定合同主体
	TenantName    string             `json:"tenantName,omitempty"`
	DeleteOldData *bool              `json:"deleteOldData,omitempty"`
	Watermarks    []*model.Watermark `json:"watermarks,omitempty"`
}

func (obj ContractAddwartermarkRequest) GetUrl() string {
	return "/v2/contract/addwartermark"
}

func (obj ContractAddwartermarkRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
