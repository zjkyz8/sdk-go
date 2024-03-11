package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
)

type SubCompanyInviteRequest struct {
	// 该公司必须为已经完成认证的契约锁公司
	CompanyName string `json:"companyName,omitempty"`
	RegisterNo  string `json:"registerNo,omitempty"`
	BizId       string `json:"bizId,omitempty"`
	CallbackUrl string `json:"callbackUrl,omitempty"`
}

func (obj SubCompanyInviteRequest) GetUrl() string {
	return "/v2/subcompany/invite"
}

func (obj SubCompanyInviteRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
