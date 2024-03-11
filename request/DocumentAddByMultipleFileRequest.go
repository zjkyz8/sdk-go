package request

import (
	"qiyuesuo/sdk/http"
)

type DocumentAddByMultipleFileRequest struct {
	// 合同ID，合同ID与业务ID二选一，不能同时为空
	ContractId string `json:"contractId,omitempty"`
	// 业务ID，合同ID与业务ID二选一，不能同时为空
	BizId string `json:"bizId,omitempty"`
	// 子公司名称，若使用业务ID添加合同文件，且合同是以子公司身份创建的，则需要传递该值，用于确定合同主体
	TenantName string `json:"tenantName,omitempty"`
	// 文件名称
	Title string `json:"title,omitempty"`
	// 待合并的文件集合 ,总文件大小：<=30M
	Files []*http.FileItem `json:"files,omitempty"`
	// 文档指定排序
	DocumentSort *int `json:"documentSort,omitempty"`
}

func (obj DocumentAddByMultipleFileRequest) GetUrl() string {
	return "/v2/document/addbyfiles"
}

func (obj DocumentAddByMultipleFileRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	parameter.AddParam("contractId", obj.ContractId)
	parameter.AddParam("bizId", obj.BizId)
	parameter.AddParam("tenantName", obj.TenantName)
	parameter.AddParam("title", obj.Title)
	parameter.AddParam("documentSort", obj.DocumentSort)
	parameter.AddListFiles("files", obj.Files)
	return parameter
}
