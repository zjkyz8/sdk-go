package model

type SealImageInfo struct {
	// UNIVERSAL_SEAL（公章），SPECIAL_SEAL（专用章），YOUTH_LEAGUE_SEAL（共青团章），FOREIGN_SEAL（外资企业公章-中文），FOREIGN_MIX_SEAL（外资企业公章-中英文混合）
	Style string `json:"style,omitempty"`
	// CIRCULAR_42（圆形、直径42mm），CIRCULAR_40（圆形、直径40mm），CIRCULAR_44（圆形、直径44mm），CIRCULAR_45（圆形、直径45mm），OVAL_45_30（椭圆、长短轴45mm/30mm）
	Spec string `json:"spec,omitempty"`
	// 用于设置分公司公章
	BranchName string `json:"branchName,omitempty"`
	// 除FOREIGN_MIX_SEAL外，其他样式均可设置
	Foot string `json:"foot,omitempty"`
	// 为13到15位的数字，生成公章时设置
	EnterpriseCode string `json:"enterpriseCode,omitempty"`
	// 用于FOREIGN_MIX_SEAL样式的印章
	EnContent string `json:"enContent,omitempty"`
	// YOUTH_LEAGUE_SEAL样式的章可设置值为1.5/1.2，此样式默认为1.2
	EdgeWidth string `json:"edgeWidth,omitempty"`
}
