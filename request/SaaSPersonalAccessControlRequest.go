package request

import (
	"encoding/json"

	"github.com/zjkyz8/sdk-go/http"
	"github.com/zjkyz8/sdk-go/model"
)

type SaaSPersonalAccessControlRequest struct {
	User    *model.User `json:"user,omitempty"`
	Operate string      `json:"operate,omitempty"`
}

func (obj SaaSPersonalAccessControlRequest) GetUrl() string {
	return "/saas/v2/personal/accesscontrol"
}

func (obj SaaSPersonalAccessControlRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
