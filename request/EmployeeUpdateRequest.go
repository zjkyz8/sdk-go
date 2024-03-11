package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
	"qiyuesuo/sdk/model"
)

type EmployeeUpdateRequest struct {
	User   *model.User `json:"user,omitempty"`
	Number string      `json:"number,omitempty"`
	// 若修改的该员工位于子公司中，则需要传递该值，默认为平台方主公司
	TenantName string `json:"tenantName,omitempty"`
}

func (obj EmployeeUpdateRequest) GetUrl() string {
	return "/v2/employee/update"
}

func (obj EmployeeUpdateRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
