package request

import (
	"qiyuesuo/sdk/http"
)

type ContractDownloadUrlRequest struct {
	// 子公司名称，若使用业务ID下载合同，且合同是以子公司身份创建的，则需要传递该值，用于确定合同主体
	TenantName string `json:"tenantName,omitempty"`
	// 是否压缩成zip格式返回单个链接。默认不压缩，每个文件返回单独的下载链接   >如果compress=false，当下载类型为合同原文，返回值中仅返回文档id，不返回合同id，其余下载类型仅返回合同id，不返回文档id；文件名称是单个文件对应的名称、下载链接是单个文档的下载链接；  >如果compress=true，返回值中仅返回合同id，不返回文档id；文件名称是合同标题、下载链接是整个合同压缩包的下载链接
	Compress string `json:"compress,omitempty"`
	// 合同ID，合同ID与业务ID二选一，不能同时为空
	ContractId string `json:"contractId,omitempty"`
	// 业务ID，合同ID与业务ID二选一，不能同时为空
	BizId string `json:"bizId,omitempty"`
	// 下载子项目，子项目可选项：CONTRACT（\"合同原文\"）、SIGNLOG（\"签署日志\"）、ATTACHMENT（\"附件\"）、NOTARY（\"存证报告\"）、ENDSIGN_ATTACHMENT（“强制结束附件”），各子项以逗号（\",\"）相隔，如CONTRACT,SIGNLOG，默认为合同文件与签署日志
	DownloadItems []string `json:"downloadItems,omitempty"`
}

func (obj ContractDownloadUrlRequest) GetUrl() string {
	return "/v2/contract/downloadurl"
}

func (obj ContractDownloadUrlRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewGetHttpParameter()
	parameter.AddParam("tenantName", obj.TenantName)
	parameter.AddParam("compress", obj.Compress)
	parameter.AddParam("contractId", obj.ContractId)
	parameter.AddParam("bizId", obj.BizId)
	if obj.DownloadItems != nil && len(obj.DownloadItems) > 0 {
		dis := ""
		for _, v := range obj.DownloadItems {
			dis += v + ","
		}
		dis = dis[:len(dis)-1]
		parameter.AddParam("downloadItems", dis)
	}
	return parameter
}
