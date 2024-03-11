package request

import (
	"encoding/json"

	"github.com/zjkyz8/sdk-go/http"
	"github.com/zjkyz8/sdk-go/model"
)

type SaaSPersonalSignAuthUrlRequest struct {
	User              *model.User      `json:"user,omitempty"`
	AuthDeadline      string           `json:"authDeadline,omitempty"`
	Remark            string           `json:"remark,omitempty"`
	CallbackUrl       string           `json:"callbackUrl,omitempty"`
	Company           *model.Company   `json:"company,omitempty"`
	AppId             string           `json:"appId,omitempty"`
	PageStyle         *model.PageStyle `json:"pageStyle,omitempty"`
	AllowPersonalAuto *bool            `json:"allowPersonalAuto,omitempty"`
}

func (obj SaaSPersonalSignAuthUrlRequest) GetUrl() string {
	return "/saas/v2/personalsign/authurl"
}

func (obj SaaSPersonalSignAuthUrlRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
