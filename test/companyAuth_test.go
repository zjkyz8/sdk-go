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
个人认证接口
*/

func TestCompanyAuthPage(t *testing.T) {
	req := request.CompanyAuthPCPageRequest{}
	req.CompanyName = "go测试公司-1"
	user := model.User{}
	user.Name = "宋一"
	user.Contact = "10000000281"
	user.ContactType = "MOBILE"
	req.Applicant = &user

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestCompanyAuthPageWithLicence(t *testing.T) {
	req := request.CompanyAuthLicensePCPageRequest{}
	req.CompanyName = "go测试公司-1"
	user := model.User{}
	user.Name = "宋一"
	user.Contact = "10000000011"
	user.ContactType = "MOBILE"
	userBs, _ := json.Marshal(user)
	req.ApplicantInfo = string(userBs)
	file, _ := os.Open("/Users/sgf/develop/临时/go/个人认证.pdf")
	req.License = &http.FileItem{file, ""}

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestCompanyAuthResult(t *testing.T) {
	req := request.CompanyAuthResultRequest{}
	req.CompanyName = "go测试公司-1"

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}
