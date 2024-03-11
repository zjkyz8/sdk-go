package request

import (
	"qiyuesuo/sdk/http"
)

type VerifierRequest struct {
	// 验签文件，仅支持PDF文件验签
	File *http.FileItem `json:"file,omitempty"`
	// 是否获取签名外观base64格式数据，默认false
	ShowImg *bool `json:"showImg,omitempty"`
}

func (obj VerifierRequest) GetUrl() string {
	return "/v2/verify"
}

func (obj VerifierRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	parameter.AddParam("showImg", obj.ShowImg)
	parameter.AddFiles("file", obj.File)
	return parameter
}
