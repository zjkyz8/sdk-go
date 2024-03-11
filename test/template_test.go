package test

import (
	"encoding/json"
	"fmt"
	"os"
	"qiyuesuo/sdk/http"
	"qiyuesuo/sdk/model"
	"qiyuesuo/sdk/request"
	"testing"
)

/**
模板接口
*/

func TestTemplateList(t *testing.T) {
	req := request.TemplateListRequest{}
	//req.ModifyTimeStart = "2020-09-09 00:00:00"

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestTemplateDetail(t *testing.T) {
	req := request.TemplateDetailRequest{}
	req.TemplateId = "2736539399083791075"

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestTemplatePage(t *testing.T) {
	req := request.TemplatePageRequest{}
	req.TemplateId = "2736539399083791075"

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestTemplateCreateByWord(t *testing.T) {
	req := request.TemplateCreateByWordV3Request{}
	file, _ := os.Open("/Users/sgf/develop/file/文件模板三.docx")
	req.File = &http.FileItem{Src: file}
	req.Title = "go模板"
	user := model.User{}
	user.Contact = "10000000011"
	user.ContactType = "MOBILE"
	var users []*model.User
	users = append(users, &user)
	userJson, _ := json.Marshal(users)
	req.ManagersInfo = string(userJson)
	req.RangeInfo = string(userJson)
	req.Status = "ENABLED"

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestTemplateEdit(t *testing.T) {
	req := request.TemplateEditRequest{}
	req.TemplateId = "3122117878250217689"
	req.Title = "go模板1"
	user := model.User{}
	user.Contact = "10000000011"
	user.ContactType = "MOBILE"
	var users []*model.User
	users = append(users, &user)
	req.Managers = users

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestTemplateStatusManage(t *testing.T) {
	req := request.TemplateStatusManageRequest{}
	req.TemplateId = "3122117878250217689"
	req.Operate = "ENABLE"

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestTemplateRemove(t *testing.T) {
	req := request.TemplateRemoveRequest{}
	req.TemplateId = "3122117878250217689"

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestTemplateDownload(t *testing.T) {
	req := request.TemplateDownloadRequest{}
	req.TemplateId = "3122388048050978977"

	file, _ := os.OpenFile("/Users/sgf/develop/临时/go/go模板2.docx", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	defer file.Close()
	err := sdkClient.Download(req, file)
	if err != nil {
		fmt.Println("error,", err.Error())
	}
}
