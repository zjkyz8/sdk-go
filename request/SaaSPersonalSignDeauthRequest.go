package request

import (
	"encoding/json"

	"github.com/zjkyz8/sdk-go/http"
	"github.com/zjkyz8/sdk-go/model"
)

type SaaSPersonalSignDeauthRequest struct {
	User    *model.User    `json:"user,omitempty"`
	Company *model.Company `json:"company,omitempty"`
	AppId   string         `json:"appId,omitempty"`
}

func (obj SaaSPersonalSignDeauthRequest) GetUrl() string {
	return "/saas/v2/personalsign/deauthorize"
}

func (obj SaaSPersonalSignDeauthRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
