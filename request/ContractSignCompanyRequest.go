package request

import (
	"encoding/json"

	"github.com/zjkyz8/sdk-go/http"
	"github.com/zjkyz8/sdk-go/model"
)

type ContractSignCompanyRequest struct {
	Param *model.SignParam
}

func (obj ContractSignCompanyRequest) GetUrl() string {
	return "/v2/contract/companysign"
}

func (obj ContractSignCompanyRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj.Param)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
