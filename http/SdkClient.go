package http

import (
	"encoding/json"
	"io"
	"qiyuesuo/sdk/commons"
)

const (
	SDK_VERSION = "GO-3.0.2"
)

type SdkClient struct {
	ServerUrl    string
	AccessToken  string
	AccessSecret string
	SdkVersion   string
}

func NewSdkClient(serverUrl string, accessToken string, accessSecret string) *SdkClient {
	sdkClient := SdkClient{
		ServerUrl:    serverUrl,
		AccessToken:  accessToken,
		AccessSecret: accessSecret,
		SdkVersion:   SDK_VERSION,
	}
	return &sdkClient
}

func (sdk *SdkClient) Service(request SdkRequest) (res map[string]interface{}, err error) {
	bs, err := sdk.sdkRequest(request)
	res = make(map[string]interface{})
	if err != nil {
		return
	}
	err = json.Unmarshal(bs, &res)
	return
}

func (sdk *SdkClient) Download(request SdkRequest, w io.Writer) (err error) {
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

func (sdk *SdkClient) sdkRequest(request SdkRequest) (bs []byte, err error) {
	url := sdk.ServerUrl + request.GetUrl()
	parameter := request.GetHttpParameter()
	header := sdk.buildHttpHeader(parameter)
	if parameter.IsJson() {
		return DoServiceWithJson(url, parameter, header)
	}
	return DoService(url, parameter, header)
}

func (sdk *SdkClient) buildHttpHeader(httpParameter *HttpParameter) *HttpHeader {
	timestamp := commons.GetTimeStamp()
	nonce := commons.GetUUID()
	var signature = commons.GetMD5(sdk.AccessToken + sdk.AccessSecret + timestamp + nonce)
	header := HttpHeader{
		AccessToken: sdk.AccessToken,
		Timestamp:   timestamp,
		Nonce:       nonce,
		Signature:   signature,
		SdkVersion:  sdk.SdkVersion,
	}
	return &header
}
