package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
	"qiyuesuo/sdk/model"
)

type ContractSignInvalidRequest struct {
	// 合同ID与业务ID二选一，不能同时为空
	ContractId string `json:"contractId,omitempty"`
	// 合同ID与业务ID二选一，不能同时为空
	BizId string `json:"bizId,omitempty"`
	// 若使用业务ID撤回作废合同，且合同是以子公司身份创建的，则需要传递该值，用于确定合同主体
	TenantName string `json:"tenantName,omitempty"`
	// 发起方签署作废文件时指定的印章，作废合同时使用。若发起方已签署，不传默认取发起方印章ID
	SealId string `json:"sealId,omitempty"`
	// 使用自定义文件作废的合同，且未设置发起方签署位置时，可以指定签署位置
	Stampers []*model.Stamper `json:"stampers,omitempty"`
	// 默认为true；传入false时，对于传入的关键字定位的签署位置，只需定位到任意一个签署位置的关键字即可
	LocateAllStamperKeywords *bool `json:"locateAllStamperKeywords,omitempty"`
}

func (obj ContractSignInvalidRequest) GetUrl() string {
	return "/v2/contract/invalidsign"
}

func (obj ContractSignInvalidRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
