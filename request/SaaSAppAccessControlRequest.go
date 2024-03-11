package request

import (
	"encoding/json"

	"github.com/zjkyz8/sdk-go/http"
	"github.com/zjkyz8/sdk-go/model"
)

type SaaSAppAccessControlRequest struct {
	Company *model.Company `json:"company,omitempty"`
	Operate string         `json:"operate,omitempty"`
}

func (obj SaaSAppAccessControlRequest) GetUrl() string {
	return "/saas/v2/app/accesscontrol"
}

func (obj SaaSAppAccessControlRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
