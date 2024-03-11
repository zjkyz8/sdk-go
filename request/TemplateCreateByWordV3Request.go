package request

import (
	"qiyuesuo/sdk/http"
)

type TemplateCreateByWordV3Request struct {
	// 模板名称
	Title string `json:"title,omitempty"`
	// 模板文件，仅支持docx格式，小于10MB
	File *http.FileItem `json:"file,omitempty"`
	// 公司名称，指定模板所属公司，未传递时默认为平台方名称
	TenantName string `json:"tenantName,omitempty"`
	// 模板中若有重名参数，是否自动重命名重名参数，自动重命名即按照 “参数名” + “数字” 的规则命名，示例 “金额1、金额2、金额3…”；默认不自动重命名参数
	RenameRepeatParam string `json:"renameRepeatParam,omitempty"`
	// 模板管理员，默认为企业模版管理员，可设置具体公司内的人员，List<User> 的json数组字符串
	ManagersInfo string `json:"managersInfo,omitempty"`
	// 模板使用范围，默认为所有人，允许自定义设置具体的人，List<User> 的json数组字符串
	RangeInfo string `json:"rangeInfo,omitempty"`
	// 此模板使用范围是否设置为所有人，默认为false；为true且rangeInfo参数为空时，设置所有人可使用
	AllUser string `json:"allUser,omitempty"`
	// 模板状态，默认为启用， ENABLED（启用）、DISABLE（停用）
	Status string `json:"status,omitempty"`
}

func (obj TemplateCreateByWordV3Request) GetUrl() string {
	return "/v3/template/createbyword"
}

func (obj TemplateCreateByWordV3Request) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	parameter.AddParam("title", obj.Title)
	parameter.AddParam("tenantName", obj.TenantName)
	parameter.AddParam("renameRepeatParam", obj.RenameRepeatParam)
	parameter.AddParam("managersInfo", obj.ManagersInfo)
	parameter.AddParam("rangeInfo", obj.RangeInfo)
	parameter.AddParam("allUser", obj.AllUser)
	parameter.AddParam("status", obj.Status)
	parameter.AddFiles("file", obj.File)
	return parameter
}
