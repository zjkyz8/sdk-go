package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
	"qiyuesuo/sdk/model"
)

type SaaSSealSignAuthUrlRequest struct {
	SealAdmin    *model.User    `json:"sealAdmin,omitempty"`
	Company      *model.Company `json:"company,omitempty"`
	AuthDeadline string         `json:"authDeadline,omitempty"`
	Remark       string         `json:"remark,omitempty"`
	CallbackUrl  string         `json:"callbackUrl,omitempty"`
	ReturnUrl    string         `json:"returnUrl,omitempty"`
	AppId        string         `json:"appId,omitempty"`
}

func (obj SaaSSealSignAuthUrlRequest) GetUrl() string {
	return "/saas/v2/sealsign/authurl"
}

func (obj SaaSSealSignAuthUrlRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
