package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
	"qiyuesuo/sdk/model"
)

type PersonalSignDeauthRequest struct {
	User *model.User `json:"user,omitempty"`
	// 若授权给子公司，则需要传递该值，用于确定授权主体
	TenantName string `json:"tenantName,omitempty"`
}

func (obj PersonalSignDeauthRequest) GetUrl() string {
	return "/v2/personalsign/deauthorize"
}

func (obj PersonalSignDeauthRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
