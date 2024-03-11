package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
)

type SealStatusManageRequest struct {
	SealId string `json:"sealId,omitempty"`
	// ENABLE（“启用”）、DISABLE（“禁用”）
	Operate string `json:"operate,omitempty"`
}

func (obj SealStatusManageRequest) GetUrl() string {
	return "/v2/seal/status"
}

func (obj SealStatusManageRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
