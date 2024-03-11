package request

import (
	"qiyuesuo/sdk/http"
	"qiyuesuo/sdk/model"
)

type UserAuthResultRequest struct {
	User *model.User
	// 认证用户联系方式，认证Id authId和用户信息 contact、contactType两者不能同时为空
	Contact string `json:"contact,omitempty"`
	// 认证用户联系方式类型，MOBILE(\"手机\")、EMAIL(\"邮箱\")
	ContactType string `json:"contactType,omitempty"`
	// 待查询的认证请求信息，认证Id authId和用户信息 contact、contactType两者不能同时为空
	AuthId string `json:"authId,omitempty"`
}

func (obj UserAuthResultRequest) GetUrl() string {
	return "/v2/personalauth/result"
}

func (obj UserAuthResultRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewGetHttpParameter()
	if obj.User != nil {
		parameter.AddParam("contact", obj.User.Contact)
		parameter.AddParam("contactType", obj.User.ContactType)
	} else {
		parameter.AddParam("contact", obj.Contact)
		parameter.AddParam("contactType", obj.ContactType)
	}
	parameter.AddParam("authId", obj.AuthId)
	return parameter
}
