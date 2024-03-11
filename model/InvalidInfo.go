package model

type InvalidInfo struct {
	// 合同ID与业务ID二选一，不能同时为空
	ContractId string `json:"contractId,omitempty"`
	// 合同ID与业务ID二选一，不能同时为空
	BizId string `json:"bizId,omitempty"`
	// 若使用业务ID撤回作废合同，且合同是以子公司身份创建的，则需要传递该值，用于确定合同主体
	TenantName               string     `json:"tenantName,omitempty"`
	Title                    string     `json:"title,omitempty"`
	FileSuffix               string     `json:"fileSuffix,omitempty"`
	DeleteDoc                *bool      `json:"deleteDoc,omitempty"`
	InvalidReason            string     `json:"invalidReason,omitempty"`
	Stampers                 []*Stamper `json:"stampers,omitempty"`
	LocateAllStamperKeywords *bool      `json:"locateAllStamperKeywords,omitempty"`
}
