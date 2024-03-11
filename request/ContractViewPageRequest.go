package request

import (
	"qiyuesuo/sdk/http"
)

type ContractViewPageRequest struct {
	// 子公司名称，若使用业务ID获取浏览页面，且合同是以子公司身份创建的，则需要传递该值，用于确定合同主体
	TenantName string `json:"tenantName,omitempty"`
	// 合同ID，合同ID和业务ID不能同时为空
	ContractId string `json:"contractId,omitempty"`
	// 业务ID，合同ID和业务ID不能同时为空
	BizId string `json:"bizId,omitempty"`
	// 文档ID，指定浏览页面默认打开的文档
	DocumentId string `json:"documentId,omitempty"`
	// 指定页面语言，cn（中文）、en（英文）；默认以发起人语言为准，合同中未传发起人时，展示中文
	Lang string `json:"lang,omitempty"`
	// 自定义页面主题色，RGB颜色(16进制)，默认取业务系统令牌配置的主题色
	ThemeColor string `json:"themeColor,omitempty"`
}

func (obj ContractViewPageRequest) GetUrl() string {
	return "/v2/contract/viewurl"
}

func (obj ContractViewPageRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewGetHttpParameter()
	parameter.AddParam("tenantName", obj.TenantName)
	parameter.AddParam("themeColor", obj.ThemeColor)
	parameter.AddParam("contractId", obj.ContractId)
	parameter.AddParam("bizId", obj.BizId)
	parameter.AddParam("documentId", obj.DocumentId)
	parameter.AddParam("lang", obj.Lang)
	return parameter
}
