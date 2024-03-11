package request

import (
	"encoding/json"

	"github.com/zjkyz8/sdk-go/http"
	"github.com/zjkyz8/sdk-go/model"
)

type OpenSSOPrivilegeUrlRequest struct {
	Company     *model.Company   `json:"company,omitempty"`
	User        *model.User      `json:"user,omitempty"`
	SuccessUrl  string           `json:"successUrl,omitempty"`
	CallbackUrl string           `json:"callbackUrl,omitempty"`
	PageStyle   *model.PageStyle `json:"pageStyle,omitempty"`
}

func (obj OpenSSOPrivilegeUrlRequest) GetUrl() string {
	return "/v2/ssoprivilege/url"
}

func (obj OpenSSOPrivilegeUrlRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
