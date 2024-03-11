package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
	"qiyuesuo/sdk/model"
)

type PersonalSignAuthQueryRequest struct {
	User *model.User `json:"user,omitempty"`
	// 默认不传，不传则返回授权 给该业务系统的所有静默签 授权记录。也可以传入某个 指定的静默签授权状态，则 接口仅返回传入的状态的所 有静默签授权记录，可传入 的状态有生效中、已过期、 已取消
	AuthStatus string `json:"authStatus,omitempty"`
}

func (obj PersonalSignAuthQueryRequest) GetUrl() string {
	return "/v2/personalsign/query"
}

func (obj PersonalSignAuthQueryRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
