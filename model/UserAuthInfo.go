package model

type UserAuthInfo struct {
	IdCardNo string `json:"idCardNo,omitempty"`
	// 仅银行卡认证时生效
	BankNo string `json:"bankNo,omitempty"`
	// 仅银行卡认证时生效
	BankMobile string `json:"bankMobile,omitempty"`
	// 姓名-USERNAME，身份证-IDCARDNO，银行卡号-BANKNO，银行卡预留手机号-BANKMOBILE
	ModifyFields []string `json:"modifyFields,omitempty"`
}
