package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
	"qiyuesuo/sdk/model"
)

type ContractDraftRequest struct {
	// 接口返回值
	Contract *model.Contract
}

func (obj ContractDraftRequest) GetUrl() string {
	return "/v2/contract/draft"
}

func (obj ContractDraftRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj.Contract)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
