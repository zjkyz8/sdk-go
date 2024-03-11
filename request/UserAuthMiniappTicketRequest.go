package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
	"qiyuesuo/sdk/model"
)

type UserAuthMiniappTicketRequest struct {
	Mode         string      `json:"mode,omitempty"`
	User         *model.User `json:"user,omitempty"`
	PaperType    string      `json:"paperType,omitempty"`
	Username     string      `json:"username,omitempty"`
	IdCardNo     string      `json:"idCardNo,omitempty"`
	BankNo       string      `json:"bankNo,omitempty"`
	BankMobile   string      `json:"bankMobile,omitempty"`
	CallbackUrl  string      `json:"callbackUrl,omitempty"`
	OtherModes   []string    `json:"otherModes,omitempty"`
	ModifyFields []string    `json:"modifyFields,omitempty"`
}

func (obj UserAuthMiniappTicketRequest) GetUrl() string {
	return "/v2/personalauth/miniappexchange"
}

func (obj UserAuthMiniappTicketRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
