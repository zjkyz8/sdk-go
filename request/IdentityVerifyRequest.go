package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
)

type IdentityVerifyRequest struct {
	Name       string `json:"name,omitempty"`
	CardNo     string `json:"cardNo,omitempty"`
	Mobile     string `json:"mobile,omitempty"`
	BankcardNo string `json:"bankcardNo,omitempty"`
}

func (obj IdentityVerifyRequest) GetUrl() string {
	return "/v2/identity/verify"
}

func (obj IdentityVerifyRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
