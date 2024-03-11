package model

type Contract struct {
	// 接口返回值
	Id string `json:"id,omitempty"`
	// 一个合同对应一个bizId，不能重复
	BizId string `json:"bizId,omitempty"`
	// 若需以子公司身份发起合同需要传递该值，默认为对接方公司
	TenantName string `json:"tenantName,omitempty"`
	// 若业务分类中未配置文件主题按规则生成，则需要传递
	Subject     string `json:"subject,omitempty"`
	Description string `json:"description,omitempty"`
	// 可由用户传入， 也可由契约锁自动生成
	Sn string `json:"sn,omitempty"`
	// 格式为yyyy-MM-dd HH:mm:ss， 默认过期时间为业务分类中配置的时
	ExpireTime string `json:"expireTime,omitempty"`
	// 格式为yyyy-MM-dd HH:mm:ss， 传递后会重置为当天23:59:59
	EndTime string `json:"endTime,omitempty"`
	// 默认为true
	Ordinal *bool `json:"ordinal,omitempty"`
	// 发起合同后不能再进行添加文档、 指定签署位置等操作
	Send     *bool     `json:"send,omitempty"`
	Category *Category `json:"category,omitempty"`
	Creator  *User     `json:"creator,omitempty"`
	// DRAFT（草稿） RECALLED（已撤回），SIGNING（签署中），REJECTED（已退回），COMPLETE（已完成），EXPIRED（已过期），FILLING（拟定中），INVALIDING（作废中），INVALIDED（已作废）
	Status string `json:"status,omitempty"`
	// 接口返回值
	Signatories      []*Signatory `json:"signatories,omitempty"`
	SignFlowStrategy string       `json:"signFlowStrategy,omitempty"`
	// 用于文件模板的填参
	TemplateParams []*TemplateParam `json:"templateParams,omitempty"`
	// 接口返回值
	Documents          []*Document         `json:"documents,omitempty"`
	CopySendReceivers  []*CopySendReceiver `json:"copySendReceivers,omitempty"`
	Tags               string              `json:"tags,omitempty"`
	RelatedContractIds []string            `json:"relatedContractIds,omitempty"`
	// 填充参数，需在业务分类下对自定义文件表单进行配置
	CustomFields []*ContractCustomField `json:"customFields,omitempty"`
	// 会在合同回调时作为参数回调
	BusinessData string `json:"businessData,omitempty"`
	// SEND：发起时抄送；FINISH：签署完成时抄送 不传时，以业务分类及公司配置为准
	CopySendTime string `json:"copySendTime,omitempty"`
}
