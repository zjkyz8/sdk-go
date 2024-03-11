package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
	"qiyuesuo/sdk/model"
)

type RoleManagementRequest struct {
	// 用于管理子公司下的角色，不指定此参数时表示管理平台方的角色
	TenantName string `json:"tenantName,omitempty"`
	// SEAL_ADMIN（“印章管理员”），TEMPLATE_ADMIN（“模板管理员”）
	Role string `json:"role,omitempty"`
	// ADD（“添加”），REMOVE（“移除”），默认为添加
	Operate string `json:"operate,omitempty"`
	// 必须为已注册的企业员工
	Users []*model.User `json:"users,omitempty"`
}

func (obj RoleManagementRequest) GetUrl() string {
	return "/v2/role/manage"
}

func (obj RoleManagementRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
