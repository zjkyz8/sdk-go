package request

import (
	"encoding/json"

	"github.com/zjkyz8/sdk-go/http"
	"github.com/zjkyz8/sdk-go/model"
)

type DocumentAddByTemplateRequest struct {
	// 合同ID与业务ID二选一，不能同时为空
	ContractId string `json:"contractId,omitempty"`
	// 合同ID与业务ID二选一，不能同时为空
	BizId string `json:"bizId,omitempty"`
	// 若使用业务ID添加合同文件，且合同是以子公司身份创建的，则需要传递该值，用于确定合同主体
	TenantName string `json:"tenantName,omitempty"`
	Title      string `json:"title,omitempty"`
	TemplateId string `json:"templateId,omitempty"`
	// 如果是参数模板，则必填（支持格式为文本、单选框、多选框格式的参数）
	TemplateParams []*model.TemplateParam `json:"templateParams,omitempty"`
	DocumentSort   *int                   `json:"documentSort,omitempty"`
}

func (obj DocumentAddByTemplateRequest) GetUrl() string {
	return "/v2/document/addbytemplate"
}

func (obj DocumentAddByTemplateRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
