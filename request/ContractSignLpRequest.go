package request

import (
	"encoding/json"

	"github.com/zjkyz8/sdk-go/http"
	"github.com/zjkyz8/sdk-go/model"
)

type ContractSignLpRequest struct {
	Param *model.SignParam
}

func (obj ContractSignLpRequest) GetUrl() string {
	return "/v2/contract/legalpersonsign"
}

func (obj ContractSignLpRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj.Param)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
