package model

type Signatory struct {
	Id string `json:"id,omitempty"`
	// COMPANY（公司）， PERSONAL（个人）
	TenantType string `json:"tenantType,omitempty"`
	// DRAFT（草稿） RECALLED（已撤回），SIGNING（签署中），REJECTED（已退回），SIGNED（已完成），EXPIRED（已过期），FILLING（拟定中），WAITING（待签署），INVALIDING（作废中），INVALIDED（已作废）
	Status string `json:"status,omitempty"`
	// delaySet为false时必填
	TenantName string    `json:"tenantName,omitempty"`
	Receiver   *User     `json:"receiver,omitempty"`
	SerialNo   *int      `json:"serialNo,omitempty"`
	Actions    []*Action `json:"actions,omitempty"`
	// 用于指定用户签署时上传的附件
	Attachments  []*Attchment  `json:"attachments,omitempty"`
	UserAuthInfo *UserAuthInfo `json:"userAuthInfo,omitempty"`
	Category     *Category     `json:"category,omitempty"`
	// 默认为false
	DelaySet *bool `json:"delaySet,omitempty"`
}
