package model

type Category struct {
	Id string `json:"id,omitempty"`
	// 如果id为空时，使用name来确定业务分类， 需要保证name对应的业务分类唯一
	Name string `json:"name,omitempty"`
}
