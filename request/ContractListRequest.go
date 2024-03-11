package request

import (
	"qiyuesuo/sdk/http"
)

type ContractListRequest struct {
	// 合同发起的开始时间，格式为yyyy-MM-dd HH:mm:ss，该字段为空时默认取六个月前的时间节点
	PublishTimeStart string `json:"publishTimeStart	,omitempty"`
	// 子公司名称
	TenantName string `json:"tenantName,omitempty"`
	// 合同发起的结束时间，格式为yyyy-MM-dd HH:mm:ss，该字段为空时默认取当前时间节点
	PublishTimeEnd string `json:"publishTimeEnd	,omitempty"`
	// 查询开始的位置(从0开始)
	SelectOffset *int `json:"selectOffset,omitempty"`
	// 查询条数限制(默认1000)
	SelectLimit *int `json:"selectLimit,omitempty"`
	// 排序(ASC:时间升序 DESC:时间降序)
	CreateTimeOrder string `json:"createTimeOrder,omitempty"`
	// 业务分类名称
	CategoryName string `json:"categoryName,omitempty"`
	// 业务分类ID
	CategoryId *float64 `json:"categoryId,omitempty"`
	// 签署方类型，默认SPONSOR，（SPONSOR：发起方，RECEIVER：接收方，ALL:全部）
	SignatoryType string `json:"signatoryType,omitempty"`
	// 合同状态(DRAFT(草稿)，RECALLED(撤回)，SIGNING(签约中)，REJECTED(已退回)，COMPLETE(已完成)，EXPIRED(已过期)，FILLING(拟定中)，INVALIDING(作废中)，INVALIDED(已作废)，FORCE_END（强制结束）
	Status string `json:"status,omitempty"`
}

func (obj ContractListRequest) GetUrl() string {
	return "/v2/contract/list"
}

func (obj ContractListRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewGetHttpParameter()
	parameter.AddParam("publishTimeStart	", obj.PublishTimeStart)
	parameter.AddParam("tenantName", obj.TenantName)
	parameter.AddParam("publishTimeEnd	", obj.PublishTimeEnd)
	parameter.AddParam("selectOffset", obj.SelectOffset)
	parameter.AddParam("selectLimit", obj.SelectLimit)
	parameter.AddParam("createTimeOrder", obj.CreateTimeOrder)
	parameter.AddParam("categoryName", obj.CategoryName)
	parameter.AddParam("categoryId", obj.CategoryId)
	parameter.AddParam("signatoryType", obj.SignatoryType)
	parameter.AddParam("status", obj.Status)
	return parameter
}
