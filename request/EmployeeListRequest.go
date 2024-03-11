package request

import (
	"qiyuesuo/sdk/http"
)

type EmployeeListRequest struct {
	// 查询条件：员工角色，SYSTEM_ADMIN（“系统管理员”）、SEAL_ADMIN（“印章管理员”）、TEMPLATE_ADMIN（“模板管理员”）
	Role string `json:"role,omitempty"`
	// 查询条件：子公司名称，若传递了则查询子公司下的员工
	TenantName string `json:"tenantName,omitempty"`
	// 查询起始位置，默认为0
	SelectOffset *int `json:"selectOffset,omitempty"`
	// 查询列表大小，默认1000
	SelectLimit *int `json:"selectLimit,omitempty"`
}

func (obj EmployeeListRequest) GetUrl() string {
	return "/v2/employee/list"
}

func (obj EmployeeListRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewGetHttpParameter()
	parameter.AddParam("role", obj.Role)
	parameter.AddParam("tenantName", obj.TenantName)
	parameter.AddParam("selectOffset", obj.SelectOffset)
	parameter.AddParam("selectLimit", obj.SelectLimit)
	return parameter
}
