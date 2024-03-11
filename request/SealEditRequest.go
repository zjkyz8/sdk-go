package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
	"qiyuesuo/sdk/model"
)

type SealEditRequest struct {
	SealId   string `json:"sealId,omitempty"`
	SealName string `json:"sealName,omitempty"`
	// 列表中的员工必须为该公司的员工
	SealUsers []*model.User `json:"sealUsers,omitempty"`
	// ADD（添加）、REMOVE（移除），默认为添加
	Operate string `json:"operate,omitempty"`
}

func (obj SealEditRequest) GetUrl() string {
	return "/v2/seal/edit"
}

func (obj SealEditRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
