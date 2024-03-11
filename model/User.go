package model

type User struct {
	Contact string `json:"contact,omitempty"`
	// MOBILE（手机号），EMAIL（邮箱），EMPLOYEEID（员工ID），NUMBER（员工编号），BIZID（用户在对接方系统的唯一标识）
	ContactType string `json:"contactType,omitempty"`
	Name        string `json:"name,omitempty"`
}
