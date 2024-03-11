package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
	"qiyuesuo/sdk/model"
)

type ContractSendRequest struct {
	// 合同ID与业务ID二选一，不能同时为空
	ContractId string `json:"contractId,omitempty"`
	// 合同ID与业务ID二选一，不能同时为空
	BizId string `json:"bizId,omitempty"`
	// 若使用业务ID发起合同，且合同是以子公司身份创建的，则需要传递该值，用于确定合同主体
	TenantName string           `json:"tenantName,omitempty"`
	Stampers   []*model.Stamper `json:"stampers,omitempty"`
	// 签署位置规则；MUST_FORBID（必须签署不可增加），MUST_ALLOW（必须签署可增加），OPTIONAL（非必须签署）；默认取业务分类中配置的规则
	StamperRule string `json:"stamperRule,omitempty"`
	// 是否必须找出全部签署位置传入的关键字，默认为true；传入false时，对于传入的关键字定位的签署位置，只需定位到任意一个签署位置的关键字即可
	LocateAllStamperKeywords *bool `json:"locateAllStamperKeywords,omitempty"`
}

func (obj ContractSendRequest) GetUrl() string {
	return "/v2/contract/send"
}

func (obj ContractSendRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
