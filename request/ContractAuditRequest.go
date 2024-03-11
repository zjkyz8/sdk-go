package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
)

type ContractAuditRequest struct {
	// 合同ID与业务ID二选一，不能同时为空
	ContractId string `json:"contractId,omitempty"`
	// 合同ID与业务ID二选一，不能同时为空
	BizId string `json:"bizId,omitempty"`
	// 若以子公司身份签署，需要传递子公司全名
	TenantName string `json:"tenantName,omitempty"`
	Pass       *bool  `json:"pass,omitempty"`
	Comment    string `json:"comment,omitempty"`
}

func (obj ContractAuditRequest) GetUrl() string {
	return "/v2/contract/employeeaudit"
}

func (obj ContractAuditRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
