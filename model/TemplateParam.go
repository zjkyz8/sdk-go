package model

type TemplateParam struct {
	Name string `json:"name,omitempty"`
	// 单行文本字数限制：1000 当使用html模板时，合同文档参数的value值说明：   > 参数类型是单行文本时，value大小不超过1000;  > 参数类型是日期时，value格式为：yyyy-MM-dd，如：2019-06-04;  > 参数类型是身份证号，value只能是15或18位的数字或字母，如：123456789123456789;  > 参数类型是单选，value只能是单选的选项,如：val1;  > 参数类型是多选，value只能是多选的选项，多个value用逗号隔开，如：val1,val2;  > 参数类型是表格，value是一个Map（键值对）数组，每个Map对应表格的每行，Map的key和value对应列名和值，示例如下： [{\"column1\":\"1\",\"column2\":\"2\",\"column3\":\"3\",\"column4\":\"4\"},{\"column1\":\"5\",\"column2\":\"6\",\"column3\":\"7\",\"column4\":\"8\"}]表示一个2行4列的表格。  > 参数类型是图片，value是图片的base64格式加前缀「data:image/png;base64,」，其中image/png为实际的图片格式，示例如下：data:image/png;base64,/9j/4AAQSk...
	Value string `json:"value,omitempty"`
	// 默认为false;传入true时，在页面上进行合同填参时该参数不可编辑
	ReadOnly string `json:"readOnly,omitempty"`
	Required *bool  `json:"required,omitempty"`
	// 单选多选不支持设置默认值，时间格式默认值需以'yyyy-MM-dd'格式提供
	DefaultValue string `json:"defaultValue,omitempty"`
	Description  string `json:"description,omitempty"`
}
