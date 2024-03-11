package model

type Document struct {
	Id        string `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	PageCount *int   `json:"pageCount,omitempty"`
	// 格式为yyyy-MM-dd HH:mm:ss
	CreateTime string `json:"createTime,omitempty"`
}
