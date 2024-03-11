package request

import (
	"qiyuesuo/sdk/http"
)

type ContractDownloadRequest struct {
	// 子公司名称，若使用业务ID下载合同，且合同是以子公司身份创建的，则需要传递该值，用于确定合同主体
	TenantName string `json:"tenantName,omitempty"`
	// 合同ID，合同ID与业务ID二选一，不能同时为空
	ContractId string `json:"contractId,omitempty"`
	// 业务ID，合同ID与业务ID二选一，不能同时为空
	BizId string `json:"bizId,omitempty"`
	// 下载子项目，子项目可选项：CONTRACT（\"合同原文\"）、SIGNLOG（\"签署日志\"）、ATTACHMENT（\"附件\"）、NOTARY（\"存证报告\"）、ENDSIGN_ATTACHMENT（“强制结束附件”），各子项以逗号（\",\"）相隔，如：CONTRACT,SIGNLOG，默认为合同文件与签署日志
	DownloadItems []string `json:"downloadItems,omitempty"`
	// 当下载的文件为单份文件时，是否压缩；默认压缩
	NeedCompressForOneFile string `json:"needCompressForOneFile,omitempty"`
}

func (obj ContractDownloadRequest) GetUrl() string {
	return "/v2/contract/download"
}

func (obj ContractDownloadRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewGetHttpParameter()
	parameter.AddParam("tenantName", obj.TenantName)
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
	parameter.AddParam("needCompressForOneFile", obj.NeedCompressForOneFile)
	return parameter
}
