package model

type Attchment struct {
	// 接口返回值
	Id    string `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	// 默认为false
	Required *bool `json:"required,omitempty"`
	// 默认为false
	NeedSign *bool `json:"needSign,omitempty"`
}
