package model

type Company struct {
	// 不可与名称同时为空
	Id string `json:"id,omitempty"`
	// 不可与id同时为空
	Name       string `json:"name,omitempty"`
	RegisterNo string `json:"registerNo,omitempty"`
}
