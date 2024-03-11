package model

type SignParam struct {
	// 合同ID与业务ID二选一，不能同时为空
	ContractId string `json:"contractId,omitempty"`
	// 合同ID与业务ID二选一，不能同时为空
	BizId string `json:"bizId,omitempty"`
	// 若以子公司身份签署，需要传递子公司全名
	TenantName string `json:"tenantName,omitempty"`
	// 若指定了签署位置，可直接在此传递印章ID，同时若下方签署位置参数中无印章ID，会从此处获取
	SealId string `json:"sealId,omitempty"`
	// 默认为true；传入false时，对于传入的关键字定位的签署位置，只需定位到任意一个签署位置的关键字即可
	LocateAllStamperKeywords *bool      `json:"locateAllStamperKeywords,omitempty"`
	Stampers                 []*Stamper `json:"stampers,omitempty"`
}
