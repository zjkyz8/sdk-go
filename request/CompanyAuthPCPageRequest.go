package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
	"qiyuesuo/sdk/model"
)

type CompanyAuthPCPageRequest struct {
	CompanyName string      `json:"companyName,omitempty"`
	Applicant   *model.User `json:"applicant,omitempty"`
	RegisterNo  string      `json:"registerNo,omitempty"`
	LegalPerson string      `json:"legalPerson,omitempty"`
	CallbackUrl string      `json:"callbackUrl,omitempty"`
	// 默认为true
	CloseButton *bool `json:"closeButton,omitempty"`
	// 支持传入字段：corpName（待认证公司名称）、registerNo（待认证公司注册号）、legalPerson（待认证公司法人姓名），传入的字段，用户在企业认证页面可支持修改，多个字段用逗号隔开（例：corpName,registerNo）
	ModifyFields string `json:"modifyFields,omitempty"`
	// cn（中文）、en（英文）；默认为企业认证申请人在契约锁设置的语言
	Lang      string           `json:"lang,omitempty"`
	PageStyle *model.PageStyle `json:"pageStyle,omitempty"`
}

func (obj CompanyAuthPCPageRequest) GetUrl() string {
	return "/companyauth/pcpage"
}

func (obj CompanyAuthPCPageRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
