package http

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"

	"github.com/zjkyz8/sdk-go/commons"
)

func DoService(interfaceUrl string, httpParameter *HttpParameter, header *HttpHeader) (bs []byte, err error) {
	method := httpParameter.HttpMethod()
	switch method {
	case http.MethodPost:
		return Post(interfaceUrl, httpParameter, header)
	case http.MethodGet:
		return Get(interfaceUrl, httpParameter, header)
	}
	return
}

// post请求json格式
func DoServiceWithJson(interfaceUrl string, httpParameter *HttpParameter, header *HttpHeader) (bs []byte, err error) {
	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, interfaceUrl, bytes.NewReader([]byte(httpParameter.jsonParamer)))
	if err != nil {
		return
	}
	// 设置header参数
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	SetHeaders(req, header)
	bs, err = DoHttpRequest(req)
	return
}

// get请求
func Get(interfaceUrl string, httpParameter *HttpParameter, header *HttpHeader) (bs []byte, err error) {
	var req *http.Request
	if len(httpParameter.Params()) > 0 {
		params := encodeParams(httpParameter.Params())
		encodedURL := url.PathEscape(params)
		interfaceUrl += "?" + encodedURL
	}
	req, err = http.NewRequest(http.MethodGet, interfaceUrl, nil)
	if err != nil {
		return
	}
	SetHeaders(req, header)
	bs, err = DoHttpRequest(req)
	return
}

// post请求，支持普通参数和文件类型参数
func Post(interfaceUrl string, httpParameter *HttpParameter, header *HttpHeader) (bs []byte, err error) {
	var req *http.Request
	if httpParameter.IsMultipart() {
		req, err = newFileUploadRequest(interfaceUrl, httpParameter, header)
	} else {
		params := encodeParams(httpParameter.Params())
		req, err = http.NewRequest(http.MethodPost, interfaceUrl, strings.NewReader(params))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		SetHeaders(req, header)
	}
	if err != nil {
		return
	}
	bs, err = DoHttpRequest(req)
	return
}

func DoDownloadWithJson(url string, httpParameter *HttpParameter, header *HttpHeader, w io.Writer) (err error) {
	var req *http.Request
	// 构造请求对象
	req, err = http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(httpParameter.jsonParamer)))
	if err != nil {
		return
	}
	// 设置header参数
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	SetHeaders(req, header)
	err = doHttpDownloadRequest(req, w)
	return
}

func DoDownload(url string, httpParameter *HttpParameter, header *HttpHeader, w io.Writer) (err error) {
	var req *http.Request
	if len(httpParameter.Params()) > 0 {
		url += "?" + encodeParams(httpParameter.Params())
	}
	req, err = http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}
	SetHeaders(req, header)
	err = doHttpDownloadRequest(req, w)
	return
}

// 创建文件上传请求
func newFileUploadRequest(url string, httpParameter *HttpParameter, header *HttpHeader) (*http.Request, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	// 普通参数
	if len(httpParameter.params) > 0 {
		for k, v := range httpParameter.Params() {
			if commons.IsNil(v) {
				continue
			}
			switch t := v.(type) {
			case string:
				if len(t) != 0 {
					writer.WriteField(k, t)
				}
			case *int, *int32, *int64:
				ele := reflect.ValueOf(v).Elem().Int()
				writer.WriteField(k, strconv.FormatInt(ele, 10))
			//case *int32:
			//	writer.WriteField(k, fmt.Sprintf("%d", *t))
			//case *int64:
			//	writer.WriteField(k, fmt.Sprintf("%d", *t))
			case *float32:
				writer.WriteField(k, fmt.Sprintf("%f", *t))
			case *float64:
				writer.WriteField(k, fmt.Sprintf("%f", *t))
			case *bool:
				if f, ok := v.(*bool); ok {
					writer.WriteField(k, fmt.Sprintf("%t", *f))
				}
			default:
				writer.WriteField(k, fmt.Sprintf("%v", t))
			}
		}
	}
	// 文件参数
	if len(httpParameter.files) > 0 {
		for k, file := range httpParameter.Files() {
			defer file.Close()
			part, err := writer.CreateFormFile(k, filepath.Base(file.GetName()))
			if err != nil {
				log.Println(err)
			}
			_, err = io.Copy(part, file.Src)
			if err != nil {
				log.Println(err)
			}
		}
	}
	// 文件列表参数
	if len(httpParameter.listFiles) > 0 {
		for k, files := range httpParameter.ListFiles() {
			if len(files) > 0 {
				for _, file := range files {
					defer file.Close()
					part, err := writer.CreateFormFile(k, filepath.Base(file.GetName()))
					if err != nil {
						log.Println(err)
					}
					_, err = io.Copy(part, file.Src)
					if err != nil {
						log.Println(err)
					}
				}
			}
		}
	}
	err := writer.Close()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, url, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	SetHeaders(req, header)
	return req, err
}

// DoHttpRequest 非下载请求
func DoHttpRequest(req *http.Request) (bs []byte, err error) {
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return
	}
	defer response.Body.Close()
	// 处理响应
	bs, err = io.ReadAll(response.Body)
	if err != nil {
		return
	}
	if response.StatusCode >= 400 {
		errMsg := "request error,httpStatus:" + response.Status
		if bs != nil && len(bs) > 0 {
			errMsg += "," + string(bs)
		}
		err = errors.New(errMsg)
		return
	}
	return
}

// doHttpDownloadRequest 下载请求
func doHttpDownloadRequest(req *http.Request, w io.Writer) (err error) {
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return
	}
	defer response.Body.Close()
	// 处理响应
	if response.StatusCode >= 400 {
		errMsg := "request error,httpStatus:" + response.Status
		bs, _ := io.ReadAll(response.Body)
		if bs != nil && len(bs) > 0 {
			errMsg += "," + string(bs)
		}
		err = errors.New(errMsg)
		return
	}
	_, err = io.Copy(w, response.Body)
	return
}

// 编码请求参数
func encodeParams(params map[string]interface{}) string {
	var parts []string
	for k, v := range params {
		if commons.IsNil(v) {
			continue
		}
		switch t := v.(type) {
		case string:
			if len(t) != 0 {
				parts = append(parts, k+commons.EQUALSIGN+fmt.Sprintf("%v", t))
			}
		case *int, *int32, *int64:
			ele := reflect.ValueOf(v).Elem().Int()
			parts = append(parts, k+commons.EQUALSIGN+strconv.FormatInt(ele, 10))
		//case *int:
		//	parts = append(parts, k+commons.EQUALSIGN+fmt.Sprintf("%d", *t))
		//case *int32:
		//	parts = append(parts, k+commons.EQUALSIGN+fmt.Sprintf("%d", *t))
		//case *int64:
		//	parts = append(parts, k+commons.EQUALSIGN+fmt.Sprintf("%d", *t))
		case *float32:
			parts = append(parts, k+commons.EQUALSIGN+fmt.Sprintf("%f", *t))
		case *float64:
			parts = append(parts, k+commons.EQUALSIGN+fmt.Sprintf("%f", *t))
		case *bool:
			if f, ok := v.(*bool); ok {
				parts = append(parts, k+commons.EQUALSIGN+fmt.Sprintf("%t", *f))
			}
		default:
			parts = append(parts, k+commons.EQUALSIGN+fmt.Sprintf("%v", t))
		}
	}
	return strings.Join(parts, commons.CONNECTOR)
}

func SetHeaders(req *http.Request, header *HttpHeader) {
	if header.AccessToken != "" {
		req.Header.Set(commons.ACCESS_TOKEN, header.AccessToken)
	}
	if header.Timestamp != "" {
		req.Header.Set(commons.TIMESTAMP, header.Timestamp)
	}
	if header.Signature != "" {
		req.Header.Set(commons.SIGNATURE, header.Signature)
	}
	if header.Nonce != "" {
		req.Header.Set(commons.NONCE, header.Nonce)
	}
	if header.SdkVersion != "" {
		req.Header.Set(commons.SDK_VERSION, header.SdkVersion)
	}
	if header.AgentToken != "" {
		req.Header.Set(commons.AGENT_TOKEN, header.AgentToken)
	}
	if header.AgentSignature != "" {
		req.Header.Set(commons.AGENT_SIGNATURE, header.AgentSignature)
	}
}
