package request

import (
	"encoding/json"

	"github.com/zjkyz8/sdk-go/http"
	"github.com/zjkyz8/sdk-go/model"
)

type SealCreatebyImageRequest struct {
	SealName    string        `json:"sealName,omitempty"`
	CallbackUrl string        `json:"callbackUrl,omitempty"`
	SealSpec    string        `json:"sealSpec,omitempty"`
	SealImage   string        `json:"sealImage,omitempty"`
	TenantName  string        `json:"tenantName,omitempty"`
	SealType    string        `json:"sealType,omitempty"`
	LpLetter    string        `json:"lpLetter,omitempty"`
	SealUsers   []*model.User `json:"sealUsers,omitempty"`
}

func (obj SealCreatebyImageRequest) GetUrl() string {
	return "/v2/seal/createbyimage"
}

func (obj SealCreatebyImageRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
