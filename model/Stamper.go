package model

type Stamper struct {
	// 公司的签署位置必须（创建合同草稿接口的返回值）
	ActionId string `json:"actionId,omitempty"`
	// 个人的签署位置必传（创建合同草稿接口的返回值）
	SignatoryId string `json:"signatoryId,omitempty"`
	// COMPANY（公章），LP（法人章），PERSONAL（个人签名），TIMESTAMP（时间戳），ACROSS_PAGE（骑缝章）
	Type_      string `json:"type,omitempty"`
	DocumentId string `json:"documentId,omitempty"`
	SealId     string `json:"sealId,omitempty"`
	Keyword    string `json:"keyword,omitempty"`
	// 1代表第1个关键字，0代表所有关键字，-1代表倒数第1个关键字；默认为1
	KeywordIndex *int `json:"keywordIndex,omitempty"`
	// 0代表所有,-1代表最后一页
	Page    *int     `json:"page,omitempty"`
	OffsetX *float64 `json:"offsetX,omitempty"`
	OffsetY *float64 `json:"offsetY,omitempty"`
	// HYPHEN（yyyy-mm-dd），Chinese（yyyy年mm月dd日（阿拉伯数字）），ALL_Chinese（yyyy年mm月dd日（中文））；仅type为TIMESTAMP时生效，默认为HYPHEN
	DatePattern string `json:"datePattern,omitempty"`
}
