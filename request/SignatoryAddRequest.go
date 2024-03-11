package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
	"qiyuesuo/sdk/model"
)

type SignatoryAddRequest struct {
	// 合同ID与业务ID二选一，不能同时为空
	ContractId string `json:"contractId,omitempty"`
	// 合同ID与业务ID二选一，不能同时为空
	BizId string `json:"bizId,omitempty"`
	// 若使用业务ID发起合同，且合同是以子公司身份创建的，则需要传递该值，用于确定合同主体
	TenantName string           `json:"tenantName,omitempty"`
	Signatory  *model.Signatory `json:"signatory,omitempty"`
	// 默认为true；传入false时，对于传入的关键字定位的签署位置，只需定位到任意一个签署位置的关键字即可
	LocateAllStamperKeywords *bool `json:"locateAllStamperKeywords,omitempty"`
}

func (obj SignatoryAddRequest) GetUrl() string {
	return "/v2/signatory/add"
}

func (obj SignatoryAddRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
