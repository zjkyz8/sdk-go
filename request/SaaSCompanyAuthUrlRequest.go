package request

import (
	"encoding/json"

	"github.com/zjkyz8/sdk-go/http"
)

type SaaSCompanyAuthUrlRequest struct {
	CompanyId string `json:"companyId,omitempty"`
}

func (obj SaaSCompanyAuthUrlRequest) GetUrl() string {
	return "/saas/v2/companyauth/url"
}

func (obj SaaSCompanyAuthUrlRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
