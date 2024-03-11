package request

import (
	"encoding/json"

	"github.com/zjkyz8/sdk-go/http"
	"github.com/zjkyz8/sdk-go/model"
)

type ContractMiniappTicketRequest struct {
	ContractId           string      `json:"contractId,omitempty"`
	BizId                string      `json:"bizId,omitempty"`
	TenantName           string      `json:"tenantName,omitempty"`
	User                 *model.User `json:"user,omitempty"`
	HidePasswordSettings *bool       `json:"hidePasswordSettings,omitempty"`
}

func (obj ContractMiniappTicketRequest) GetUrl() string {
	return "/v2/contract/miniappexchange"
}

func (obj ContractMiniappTicketRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
