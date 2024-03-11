package request

import (
	"qiyuesuo/sdk/http"
)

type TemplateListRequest struct {
	// 查询条件：子公司名称，若传递了则查询子公司下的模板，未传递时默认为平台方名称
	TenantName string `json:"tenantName,omitempty"`
	// 查询起始位置，默认为0
	SelectOffset string `json:"selectOffset,omitempty"`
	// 查询列表大小，默认1000
	SelectLimit string `json:"selectLimit,omitempty"`
	// 数据变更（创建、变更）时间的起始时间，格式为yyyy-MM-dd HH:mm:ss，默认无限制
	ModifyTimeStart string `json:"modifyTimeStart,omitempty"`
	// 数据变更（创建、更新）时间的结束时间，格式为yyyy-MM-dd HH:mm:ss，默认无限制
	ModifyTimeEnd string `json:"modifyTimeEnd,omitempty"`
}

func (obj TemplateListRequest) GetUrl() string {
	return "/v2/template/list"
}

func (obj TemplateListRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewGetHttpParameter()
	parameter.AddParam("tenantName", obj.TenantName)
	parameter.AddParam("selectOffset", obj.SelectOffset)
	parameter.AddParam("selectLimit", obj.SelectLimit)
	parameter.AddParam("modifyTimeStart", obj.ModifyTimeStart)
	parameter.AddParam("modifyTimeEnd", obj.ModifyTimeEnd)
	return parameter
}
