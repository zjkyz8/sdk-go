package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
	"qiyuesuo/sdk/model"
)

type UserAuthPageRequest struct {
	// IVS（三要素）、FACE（人脸认证）、BANK（银行卡认证），默认为DEFAULT（不指定认证模式，默认使用可使用的所有认证）
	Mode         string      `json:"mode,omitempty"`
	User         *model.User `json:"user,omitempty"`
	PaperType    string      `json:"paperType,omitempty"`
	Username     string      `json:"username,omitempty"`
	IdCardNo     string      `json:"idCardNo,omitempty"`
	BankNo       string      `json:"bankNo,omitempty"`
	BankMobile   string      `json:"bankMobile,omitempty"`
	CallbackUrl  string      `json:"callbackUrl,omitempty"`
	CallbackPage string      `json:"callbackPage,omitempty"`
	// 三要素-IVS，人脸认证-FACE，银行卡认证-BANK，人工审核-MANUAL
	OtherModes []string `json:"otherModes,omitempty"`
	// 姓名-USERNAME，身份证-IDCARDNO，银行卡号-BANKNO，银行卡预留手机号-BANKMOBILE
	ModifyFields []string `json:"modifyFields,omitempty"`
	// cn（中文）、en（英文）；默认为当前认证人在契约锁设置的语言
	Lang      string           `json:"lang,omitempty"`
	PageStyle *model.PageStyle `json:"pageStyle,omitempty"`
}

func (obj UserAuthPageRequest) GetUrl() string {
	return "/v2/personalauth"
}

func (obj UserAuthPageRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
