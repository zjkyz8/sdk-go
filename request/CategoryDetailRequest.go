package request

import (
	"qiyuesuo/sdk/http"
)

type CategoryDetailRequest struct {
	// 子公司名称，若使用业务分类名称获取业务分类详情，且业务分类是以子公司身份创建的，则需要传递该值，用于确定业务分类主体
	TenantName string `json:"tenantName,omitempty"`
	// 业务分类名称，业务分类ID和业务分类名称不能同时为空
	CategoryName string `json:"categoryName,omitempty"`
	// 业务分类ID，业务分类ID和业务分类名称不能同时为空
	CategoryId string `json:"categoryId,omitempty"`
}

func (obj CategoryDetailRequest) GetUrl() string {
	return "/v2/category/detail"
}

func (obj CategoryDetailRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewGetHttpParameter()
	parameter.AddParam("tenantName", obj.TenantName)
	parameter.AddParam("categoryName", obj.CategoryName)
	parameter.AddParam("categoryId", obj.CategoryId)
	return parameter
}
