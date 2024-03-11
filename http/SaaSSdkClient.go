package http

import (
	"encoding/json"
	"io"

	"github.com/zjkyz8/sdk-go/commons"
)

type SaaSSdkClient struct {
	ServerUrl    string
	AccessToken  string
	AccessSecret string
	SdkVersion   string
	AgentToken   string
	AgentSecret  string
}

func NewSaaSSdkClient(serverUrl string, agentToken string, agentSecret string) *SaaSSdkClient {
	SaaSSdkClient := SaaSSdkClient{
		ServerUrl:   serverUrl,
		AgentToken:  agentToken,
		AgentSecret: agentSecret,
		SdkVersion:  SDK_VERSION,
	}
	return &SaaSSdkClient
}

func (sdk *SaaSSdkClient) Service(request SdkRequest) (res map[string]interface{}, err error) {
	bs, err := sdk.sdkRequest(request)
	res = make(map[string]interface{})
	if err != nil {
		return
	}
	err = json.Unmarshal(bs, &res)
	return
}

func (sdk *SaaSSdkClient) Download(request SdkRequest, w io.Writer) (err error) {
	url := sdk.ServerUrl + request.GetUrl()
	parameter := request.GetHttpParameter()
	header := sdk.buildHttpHeader(parameter)
	if parameter.IsJson() {
		err = DoDownloadWithJson(url, parameter, header, w)
	} else {
		err = DoDownload(url, parameter, header, w)
	}
	return
}

func (sdk *SaaSSdkClient) sdkRequest(request SdkRequest) (bs []byte, err error) {
	url := sdk.ServerUrl + request.GetUrl()
	parameter := request.GetHttpParameter()
	header := sdk.buildHttpHeader(parameter)
	if parameter.IsJson() {
		return DoServiceWithJson(url, parameter, header)
	}
	return DoService(url, parameter, header)
}

func (client *SaaSSdkClient) buildHttpHeader(httpParameter *HttpParameter) *HttpHeader {
	timestamp := commons.GetTimeStamp()
	nonce := commons.GetUUID()
	agentSignature := commons.GetMD5(client.AgentToken + client.AgentSecret + timestamp + nonce)
	header := HttpHeader{
		AgentToken:     client.AgentToken,
		Timestamp:      timestamp,
		Nonce:          nonce,
		AgentSignature: agentSignature,
		SdkVersion:     client.SdkVersion,
	}
	if client.AccessToken != "" {
		signature := commons.GetMD5(client.AccessToken + client.AccessSecret + timestamp + nonce)
		header.AccessToken = client.AccessToken
		header.Signature = signature
	}
	return &header
}
