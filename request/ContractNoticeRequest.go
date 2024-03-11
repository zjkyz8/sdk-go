package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
)

type ContractNoticeRequest struct {
	// 合同ID与业务ID二选一，不能同时为空
	ContractId string `json:"contractId,omitempty"`
	// 合同ID与业务ID二选一，不能同时为空
	BizId string `json:"bizId,omitempty"`
	// 若使用业务ID催签，且合同是以子公司身份创建的，则需要传递该值，用于确定合同主体
	TenantName  string `json:"tenantName,omitempty"`
	SignatoryId string `json:"signatoryId,omitempty"`
	// 默认为true; signing传参false时，除签署中合同外，可催签拟定中、作废待确认的合同
	OnlySigning *bool `json:"onlySigning,omitempty"`
}

func (obj ContractNoticeRequest) GetUrl() string {
	return "/v2/contract/notice"
}

func (obj ContractNoticeRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
