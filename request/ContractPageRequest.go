package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
	"qiyuesuo/sdk/model"
)

type ContractPageRequest struct {
	// 合同ID和业务ID不能同时为空
	ContractId string `json:"contractId,omitempty"`
	// 合同ID和业务ID不能同时为空
	BizId string `json:"bizId,omitempty"`
	// 若使用业务ID获取签署页面，且合同是以子公司身份创建的，则需要传递该值，用于确定合同主体
	TenantName string `json:"tenantName,omitempty"`
	// 签署页面：SIGN，文件打印页面：PRINT，直接签署页面（H5）：DIRECT_SIGN，默认为签署页面仅支持PC端的操作：PRINT
	PageType     string      `json:"pageType,omitempty"`
	User         *model.User `json:"user,omitempty"`
	CallbackPage string      `json:"callbackPage,omitempty"`
	// 默认false允许显示返回按钮
	HideReturnButton *bool `json:"hideReturnButton,omitempty"`
	// 默认false允许显示撤回按钮 ；注：当参数user不具备该按钮使用权限时，按钮隐藏
	HideRecallButton *bool `json:"hideRecallButton,omitempty"`
	// 默认true隐藏退回按钮；注：当参数user不具备该按钮使用权限时，按钮隐藏
	HideRejectButton *bool `json:"hideRejectButton,omitempty"`
	// 默认false允许显示结束签署按钮；注：当参数user不具备该按钮使用权限时，按钮隐藏
	HideEndButton *bool `json:"hideEndButton,omitempty"`
	// 默认false允许显示签署状态按钮
	HideStateButton *bool `json:"hideStateButton,omitempty"`
	// 默认为false；为true时，用户首次在签署页面使用验证码签署成功后，后续30分钟内在同一客户端该用户可使用密码在签署页面进行合同签署（后续该参数也需要为true），30分钟内签署成功可使用密码签署的时间按照最后一次签署成功时刻重置30分钟
	AllowPasswordSign *bool `json:"allowPasswordSign,omitempty"`
	// cn（中文）、en（英文）；默认为接口中传的user在契约锁设置的语言
	Lang                 string           `json:"lang,omitempty"`
	PageStyle            *model.PageStyle `json:"pageStyle,omitempty"`
	HidePasswordSettings *bool            `json:"hidePasswordSettings,omitempty"`
}

func (obj ContractPageRequest) GetUrl() string {
	return "/v2/contract/pageurl"
}

func (obj ContractPageRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
