package http

type SdkRequest interface {
	GetUrl() string

	GetHttpParameter() *HttpParameter
}
