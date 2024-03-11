package request

import (
	"qiyuesuo/sdk/http"
)

type SaaSBindingRequest struct {
}

func (obj SaaSBindingRequest) GetUrl() string {
	return "/saas/v2/binding"
}

func (obj SaaSBindingRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewGetHttpParameter()
	return parameter
}
