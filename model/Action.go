package model

type Action struct {
	// 接口返回值
	Id string `json:"id,omitempty"`
	// COMPANY（企业签章），OPERATOR（经办人签字） LP（法定代表人签字），AUDIT（审批），PERSONAL（审批并签字）
	Type_ string `json:"type,omitempty"`
	// 接口返回值）：INIT（初始化） START（已启动），FINISH（已完成），STOP（意外终止）
	Status string `json:"status,omitempty"`
	Name   string `json:"name,omitempty"`
	// 从0开始
	SerialNo *int `json:"serialNo,omitempty"`
	// 用于指定签署时所有印章
	SealId    string  `json:"sealId,omitempty"`
	Operators []*User `json:"operators,omitempty"`
	// 用于签署时可选印章列表(若传了sealId优先拿sealId,不可超过9个)
	SealIds []string `json:"sealIds,omitempty"`
	// 默认为false不进行自动签（仅支持发起方下的企业签章动作；若该值为true时，对应的业务分类不应预设签署流程，并且参数send需要设为false，后续调用发起合同接口传入签署位置）
	AutoSign      *bool    `json:"autoSign,omitempty"`
	CorpSealIds   []string `json:"corpSealIds,omitempty"`
	CorpOperators []*User  `json:"corpOperators,omitempty"`
}
