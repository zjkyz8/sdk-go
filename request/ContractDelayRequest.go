package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
)

type ContractDelayRequest struct {
	// 合同ID与业务ID二选一，不能同时为空
	ContractId string `json:"contractId,omitempty"`
	// 合同ID与业务ID二选一，不能同时为空
	BizId string `json:"bizId,omitempty"`
	// 若使用业务ID发起合同，且合同是以子公司身份创建的，则需要传递该值，用于确定合同主体
	TenantName string `json:"tenantName,omitempty"`
	// 延期天数与更新合同过期时间二选一，不能同时为空，同时传入优先取延期天数； 新的过期时间=当前设置的过期时间+延期天数
	Days *int `json:"days,omitempty"`
	// 延期天数与更新合同过期时间二选一，不能同时为空，同时传入优先取延期天数； 格式为yyyy-MM-dd，不能早于当前设置的过期时间
	ExpireDate string `json:"expireDate,omitempty"`
}

func (obj ContractDelayRequest) GetUrl() string {
	return "/v2/contract/delay"
}

func (obj ContractDelayRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
