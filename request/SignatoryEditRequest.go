package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
	"qiyuesuo/sdk/model"
)

type SignatoryEditRequest struct {
	SignatoryId string      `json:"signatoryId,omitempty"`
	TenantName  string      `json:"tenantName,omitempty"`
	Receiver    *model.User `json:"receiver,omitempty"`
}

func (obj SignatoryEditRequest) GetUrl() string {
	return "/v2/signatory/edit"
}

func (obj SignatoryEditRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
