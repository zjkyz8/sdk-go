package request

import (
	"qiyuesuo/sdk/http"
	"qiyuesuo/sdk/model"
)

type SaaSUserAuthResultRequest struct {
	User   *model.User `json:"user,omitempty"`
	AuthId string      `json:"authId,omitempty"`
}

func (obj SaaSUserAuthResultRequest) GetUrl() string {
	return "/saas/v2/personalauth/result"
}

func (obj SaaSUserAuthResultRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewGetHttpParameter()
	if obj.User != nil {
		parameter.AddParam("contact", obj.User.Contact)
		parameter.AddParam("contactType", obj.User.ContactType)
	}
	parameter.AddParam("authId", obj.AuthId)
	return parameter
}
