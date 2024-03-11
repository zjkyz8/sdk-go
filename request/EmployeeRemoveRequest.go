package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
	"qiyuesuo/sdk/model"
)

type EmployeeRemoveRequest struct {
	Param *model.User `json:"user,omitempty"`
	// 若移除的该员工位于子公司中，则需要传递该值，默认为平台方主公司
	TenantName string `json:"tenantName,omitempty"`
	// 该用户会继承移除员工的权限，该用户必须为公司员工
	TransferEmployee string `json:"transferEmployee,omitempty"`
}

type originalEmployeeRemoveRequest struct {
	Contact string `json:"contact,omitempty"`
	// MOBILE（手机号）， EMAIL（邮箱），EMPLOYEEID（员工ID）
	ContactType string `json:"contactType,omitempty"`
	// 若移除的该员工位于子公司中，则需要传递该值，默认为平台方主公司
	TenantName string `json:"tenantName,omitempty"`
	// 该用户会继承移除员工的权限，该用户必须为公司员工
	TransferEmployee string `json:"transferEmployee,omitempty"`
}

func (obj EmployeeRemoveRequest) GetUrl() string {
	return "/v2/employee/remove"
}

func (obj EmployeeRemoveRequest) GetHttpParameter() *http.HttpParameter {
	// 多包装一层 EmployeeRemoveRequest 的原因：与其他语言的 sdk 参数保持一致，方便生成在线调试代码
	parameter := http.NewPostHttpParameter()
	originalObj := originalEmployeeRemoveRequest{}
	if obj.Param != nil {
		originalObj.Contact = obj.Param.Contact
		originalObj.ContactType = obj.Param.ContactType
	}
	originalObj.TenantName = obj.TenantName
	originalObj.TransferEmployee = obj.TransferEmployee

	jsonBytes, _ := json.Marshal(originalObj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
