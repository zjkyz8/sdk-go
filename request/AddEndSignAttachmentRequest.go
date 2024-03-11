package request

import (
	"qiyuesuo/sdk/http"
)

type AddEndSignAttachmentRequest struct {
	// 合同ID，合同ID与业务ID二选一，不能同时为空
	ContractId string `json:"contractId,omitempty"`
	// 业务ID，合同ID与业务ID二选一，不能同时为空
	BizId string `json:"bizId,omitempty"`
	// 子公司名称，若使用业务ID撤回作废合同，且合同是以子公司身份创建的，则需要传递该值，用于确定合同主体
	TenantName string `json:"tenantName,omitempty"`
	// 附件文件
	File *http.FileItem `json:"file,omitempty"`
	// 附件名称
	Title string `json:"title,omitempty"`
	// 文件类型（文件后缀）： doc，docx，pdf，jpeg，png，jpg，gif，tiff，html，htm，xls，xlsx
	FileSuffix string `json:"fileSuffix,omitempty"`
}

func (obj AddEndSignAttachmentRequest) GetUrl() string {
	return "/v2/contract/forceend/addattachment"
}

func (obj AddEndSignAttachmentRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	parameter.AddParam("contractId", obj.ContractId)
	parameter.AddParam("bizId", obj.BizId)
	parameter.AddParam("tenantName", obj.TenantName)
	parameter.AddParam("title", obj.Title)
	parameter.AddParam("fileSuffix", obj.FileSuffix)
	parameter.AddFiles("file", obj.File)
	return parameter
}
