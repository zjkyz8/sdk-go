package request

import (
	"qiyuesuo/sdk/http"
)

type SealImageRequest struct {
	// 印章ID
	SealId string `json:"sealId,omitempty"`
}

func (obj SealImageRequest) GetUrl() string {
	return "/v2/seal/image"
}

func (obj SealImageRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewGetHttpParameter()
	parameter.AddParam("sealId", obj.SealId)
	return parameter
}
