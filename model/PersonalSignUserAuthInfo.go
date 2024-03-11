package model

type PersonalSignUserAuthInfo struct {
	// IVS（三要素）、FACE（人脸认证）、ALIPAY（支付宝认证）、BANK（银行卡认证）， 默认为DEFAULT（不指定认证模式，默认使用可使用的所有认证）
	Mode string `json:"mode,omitempty"`
	// IDCARD(二代身份证),PASSPORT(护照),HKMP(港澳通行证),MTPS(台胞证)，默认为IDCARD
	PaperType string `json:"paperType,omitempty"`
	Username  string `json:"username,omitempty"`
	IdCardNo  string `json:"idCardNo,omitempty"`
	// 仅银行卡认证时生效
	BankNo string `json:"bankNo,omitempty"`
	// 仅银行卡认证时生效
	BankMobile string `json:"bankMobile,omitempty"`
	// 三要素-IVS，人脸认证-FACE，支付宝认证-ALIPAY，银行卡认证-BANK， 人工审核-MANUAL
	OtherModes []string `json:"otherModes,omitempty"`
	// 姓名-USERNAME，身份证-IDCARDNO，银行卡号-BANKNO，银行卡预留手机号-BANKMOBILE
	ModifyFields []string `json:"modifyFields,omitempty"`
}
