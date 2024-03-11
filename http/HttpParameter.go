package http

import (
	"net/http"
)

type HttpParameter struct {
	jsonParamer string
	params      map[string]interface{}
	files       map[string]*FileItem
	listFiles   map[string][]*FileItem
	httpMethod  string
}

func NewGetHttpParameter() *HttpParameter {
	param := HttpParameter{
		params:     make(map[string]interface{}),
		files:      make(map[string]*FileItem),
		listFiles:  make(map[string][]*FileItem),
		httpMethod: http.MethodGet,
	}
	return &param
}

func NewPostHttpParameter() *HttpParameter {
	param := HttpParameter{
		params:     make(map[string]interface{}),
		files:      make(map[string]*FileItem),
		listFiles:  make(map[string][]*FileItem),
		httpMethod: http.MethodPost,
	}
	return &param
}

func (h *HttpParameter) JsonParamer() string {
	return h.jsonParamer
}

func (h *HttpParameter) SetJsonParamer(jsonParamer string) {
	h.jsonParamer = jsonParamer
}

func (h *HttpParameter) Params() map[string]interface{} {
	return h.params
}

func (h *HttpParameter) AddParam(key string, value interface{}) {
	if value != nil && value != "" {
		h.params[key] = value
	}
}

func (h *HttpParameter) SetParams(params map[string]interface{}) {
	h.params = params
}

func (h *HttpParameter) Files() map[string]*FileItem {
	return h.files
}

func (h *HttpParameter) AddFiles(key string, file *FileItem) {
	if file != nil {
		h.files[key] = file
	}
}

func (h *HttpParameter) SetFiles(files map[string]*FileItem) {
	h.files = files
}

func (h *HttpParameter) ListFiles() map[string][]*FileItem {
	return h.listFiles
}

func (h *HttpParameter) AddListFiles(key string, listFiles []*FileItem) {
	if listFiles != nil {
		h.listFiles[key] = listFiles
	}
}

func (h *HttpParameter) SetListFiles(listFiles map[string][]*FileItem) {
	h.listFiles = listFiles
}

func (h *HttpParameter) HttpMethod() string {
	return h.httpMethod
}

func (h *HttpParameter) IsJson() bool {
	if len(h.jsonParamer) == 0 {
		return false
	}
	return true
}

func (h *HttpParameter) IsMultipart() bool {
	if len(h.files) == 0 && len(h.listFiles) == 0 {
		return false
	}
	return true
}
