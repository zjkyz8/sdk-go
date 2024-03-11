package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
	"qiyuesuo/sdk/model"
)

type EmployeeCreateRequest struct {
	User   *model.User `json:"user,omitempty"`
	Number string      `json:"number,omitempty"`
	// 若该员工需添加到子公司中，则需要传递该值，默认为平台方主公司
	TenantName string `json:"tenantName,omitempty"`
}

func (obj EmployeeCreateRequest) GetUrl() string {
	return "/v2/employee/create"
}

func (obj EmployeeCreateRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
