package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
	"qiyuesuo/sdk/model"
)

type EmployeeDetailRequest struct {
	User *model.User `json:"user,omitempty"`
	// 若该员工为子公司员工，则需要传递该值，默认为平台方主公司
	TenantName string `json:"tenantName,omitempty"`
}

func (obj EmployeeDetailRequest) GetUrl() string {
	return "/v2/employee/detail"
}

func (obj EmployeeDetailRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
