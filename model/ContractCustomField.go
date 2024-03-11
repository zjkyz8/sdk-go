package model

type ContractCustomField struct {
	// key和name至少一个非空
	Key string `json:"key,omitempty"`
	// key和name至少一个非空
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}
