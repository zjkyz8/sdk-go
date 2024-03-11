package request

import (
	"qiyuesuo/sdk/http"
)

type SealListRequest struct {
	// 查询条件：子公司名称，若传递了则查询子公司下的印章，默认为平台方主公司
	TenantName string `json:"tenantName,omitempty"`
	// 查询起始位置，默认为0
	SelectOffset *int `json:"selectOffset,omitempty"`
	// 查询列表大小，默认1000
	SelectLimit *int `json:"selectLimit,omitempty"`
	// 数据变更（创建、变更）时间的起始时间，格式为yyyy-MM-dd HH:mm:ss，默认无限制
	ModifyTimeStart string `json:"modifyTimeStart,omitempty"`
	// 数据变更（创建、更新）时间的结束时间，格式为yyyy-MM-dd HH:mm:ss，默认无限制
	ModifyTimeEnd string `json:"modifyTimeEnd,omitempty"`
	// 查询条件：印章状态：NORMAL（正常）、FREEZE（停用）、ALL（包含正常与停用状态）
	Status string `json:"status,omitempty"`
}

func (obj SealListRequest) GetUrl() string {
	return "/v2/seal/list"
}

func (obj SealListRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewGetHttpParameter()
	parameter.AddParam("tenantName", obj.TenantName)
	parameter.AddParam("selectOffset", obj.SelectOffset)
	parameter.AddParam("selectLimit", obj.SelectLimit)
	parameter.AddParam("modifyTimeStart", obj.ModifyTimeStart)
	parameter.AddParam("modifyTimeEnd", obj.ModifyTimeEnd)
	parameter.AddParam("status", obj.Status)
	return parameter
}
