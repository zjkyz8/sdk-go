package model

type CopySendReceiver struct {
	Name               string `json:"name,omitempty"`
	Receiver           *User  `json:"receiver,omitempty"`
	ProhibitType       string `json:"prohibitType,omitempty"`
	CopySendPersonType string `json:"copySendPersonType,omitempty"`
	// 格式为yyyy-MM-dd HH:mm:ss
	CreateTime string `json:"createTime,omitempty"`
	Send       string `json:"send,omitempty"`
}
