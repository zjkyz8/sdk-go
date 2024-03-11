package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
	"qiyuesuo/sdk/model"
)

type SaaSUserAuthPageRequest struct {
	Mode         string           `json:"mode,omitempty"`
	User         *model.User      `json:"user,omitempty"`
	PaperType    string           `json:"paperType,omitempty"`
	Username     string           `json:"username,omitempty"`
	IdCardNo     string           `json:"idCardNo,omitempty"`
	BankNo       string           `json:"bankNo,omitempty"`
	BankMobile   string           `json:"bankMobile,omitempty"`
	CallbackUrl  string           `json:"callbackUrl,omitempty"`
	CallbackPage string           `json:"callbackPage,omitempty"`
	ManualSwitch *bool            `json:"manualSwitch,omitempty"`
	OtherModes   []string         `json:"otherModes,omitempty"`
	ModifyFields []string         `json:"modifyFields,omitempty"`
	Lang         string           `json:"lang,omitempty"`
	PageStyle    *model.PageStyle `json:"pageStyle,omitempty"`
}

func (obj SaaSUserAuthPageRequest) GetUrl() string {
	return "/saas/v2/personalauth"
}

func (obj SaaSUserAuthPageRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
