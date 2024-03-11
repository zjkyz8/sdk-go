package request

import (
	"github.com/zjkyz8/sdk-go/http"
)

type CompanyAuthLicensePCPageRequest struct {
	// 待认证公司名称
	CompanyName string `json:"companyName,omitempty"`
	// 认证提交人信息（申请者姓名name， 联系方式contact，联系方式类型contactType：MOBILE、EMAIL），企业认证通过后，认证提交人会自动成为该企业的系统管理员(例：{\"name\":\"aaa\",\"contact\": \"15100000000\",\"contactType\": \"MOBILE\"})
	ApplicantInfo string `json:"applicantInfo,omitempty"`
	// 待认证公司注册号
	RegisterNo string `json:"registerNo,omitempty"`
	// 待认证公司法人姓名
	LegalPerson string `json:"legalPerson,omitempty"`
	// 认证回调地址
	CallbackUrl string `json:"callbackUrl,omitempty"`
	// 是否显示关闭按钮，默认为true
	CloseButton string `json:"closeButton,omitempty"`
	// 营业执照
	License *http.FileItem `json:"license,omitempty"`
	// 可修改项参数，支持传入字段：corpName（待认证公司名称）、registerNo（待认证公司注册号）、legalPerson（待认证公司法人姓名）、license（营业执照），传入的字段，用户在企业认证页面可支持修改，多个字段用逗号隔开（例：corpName,registerNo）
	ModifyFields string `json:"modifyFields,omitempty"`
	// 指定页面语言，cn（中文）、en（英文）；默认为企业认证申请人在契约锁设置的语言
	Lang string `json:"lang,omitempty"`
	// 页面风格，PageStyle JSON字符串
	PageStyleInfo string `json:"pageStyleInfo,omitempty"`
}

func (obj CompanyAuthLicensePCPageRequest) GetUrl() string {
	return "/companyauth/pcpagewithlicense"
}

func (obj CompanyAuthLicensePCPageRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	parameter.AddParam("companyName", obj.CompanyName)
	parameter.AddParam("applicantInfo", obj.ApplicantInfo)
	parameter.AddParam("registerNo", obj.RegisterNo)
	parameter.AddParam("legalPerson", obj.LegalPerson)
	parameter.AddParam("callbackUrl", obj.CallbackUrl)
	parameter.AddParam("closeButton", obj.CloseButton)
	parameter.AddParam("modifyFields", obj.ModifyFields)
	parameter.AddParam("lang", obj.Lang)
	parameter.AddParam("pageStyleInfo", obj.PageStyleInfo)
	parameter.AddFiles("license", obj.License)
	return parameter
}
