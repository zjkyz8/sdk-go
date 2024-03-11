package request

import (
	"qiyuesuo/sdk/http"
)

type SubCompanyListRequest struct {
}

func (obj SubCompanyListRequest) GetUrl() string {
	return "/v2/company/list"
}

func (obj SubCompanyListRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewGetHttpParameter()
	return parameter
}
