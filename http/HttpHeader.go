package http

type HttpHeader struct {
	ContentType    string
	AccessToken    string
	Timestamp      string
	Nonce          string
	Signature      string
	SdkVersion     string
	AgentToken     string
	AgentSignature string
}
