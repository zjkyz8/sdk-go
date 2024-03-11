package request

import (
	"qiyuesuo/sdk/http"
)

type SealRemoveRequest struct {
	// 印章ID
	SealId string `json:"sealId,omitempty"`
}

func (obj SealRemoveRequest) GetUrl() string {
	return "/v2/seal/remove"
}

func (obj SealRemoveRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewGetHttpParameter()
	parameter.AddParam("sealId", obj.SealId)
	return parameter
}
