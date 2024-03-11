package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
	"qiyuesuo/sdk/model"
)

type ContractInvalidByFileRequest struct {
	// 作废相关信息invalidContractBean的json字符串
	InvalidInfo *model.InvalidInfo `json:"invalidInfo,omitempty"`
	// 自定义作废文件
	File *http.FileItem `json:"file,omitempty"`
	// 作废通知接收人， 抄送人列表List<CopySendReceiver> 的json数组字符串，仅单方作废时生效；传入的参数里联系方式重复，以第一条为准
	ReceiversInfo []*model.CopySendReceiver
	Receivers     string `json:"receivers,omitempty"`
}

func (obj ContractInvalidByFileRequest) GetUrl() string {
	return "/v2/contract/invalidbyfile"
}

func (obj ContractInvalidByFileRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	if obj.InvalidInfo != nil {
		jsonBytes, _ := json.Marshal(obj.InvalidInfo)
		parameter.AddParam("invalidInfo", string(jsonBytes))
	}
	if obj.ReceiversInfo != nil {
		jsonBytes, _ := json.Marshal(obj.ReceiversInfo)
		parameter.AddParam("receivers", string(jsonBytes))
	} else {
		parameter.AddParam("receivers", obj.Receivers)
	}
	parameter.AddFiles("file", obj.File)
	return parameter
}
