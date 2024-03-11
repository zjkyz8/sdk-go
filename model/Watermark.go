package model

type Watermark struct {
	// NO（合同编号水印）、TEXT（文本水印）、QRCODE（合同二维码水印）、PIC（图片水印）、GRATING（光栅水印）、NOTARY（存证二维码水印）
	Type_ string `json:"type,omitempty"`
	// 当类型是TEXT/PIC/GRATING时必传；类型是TEXT时，要求文本不超过20个字符；类型是PIC时，要求图片大小不大于200Kb，水印内容为图片的base64格式；类型是GRATING时，要求内容只能是一位数字或字母
	Content string `json:"content,omitempty"`
	// 默认是16；取值范围为6到72
	FontSize *float64 `json:"fontSize,omitempty"`
	// 默认为0；取值范围为0到180
	Angle *int `json:"angle,omitempty"`
	// 默认为#000000（黑色）；请在公有云网站上选取
	FontColor string `json:"fontColor,omitempty"`
	// 默认为左上，取值范围：TOP_LEFT（左上角）、TOP_RIGHT（右上角）、BOTTOM_LEFT（左下角）、BOTTOM_RIGHT（右下角）、MIDDLE_CENTER（居中）、TILE（平铺）、FILL（填充）；其中光栅水印和合同二维码只能选择前四种
	Position string `json:"position,omitempty"`
	// 默认为30（即30%），取值范围为0到100
	Diaphaneity *float64 `json:"diaphaneity,omitempty"`
	// 默认为0.5（即20mm*20mm）；取值范围为0.1到4
	ImgScale *float64 `json:"imgScale,omitempty"`
	// 默认MODERATE（适中），取值范围：SPARSE（稀疏）、MODERATE（适中）、DENSE（密集）
	TileDensity string `json:"tileDensity,omitempty"`
	// 默认SQUARE_10
	GratingSize string `json:"gratingSize,omitempty"`
	// 水印类型是合同二维码时，二维码的查看权限，默认为PRIVATE（合同参与人登录查看），取值范围：PRIVATE（合同参与人登录查看）、PUBLIC（所有人均可访问）
	ViewPermission string `json:"viewPermission,omitempty"`
}
