package test

import (
	"fmt"
	"qiyuesuo/sdk/model"
	"qiyuesuo/sdk/request"
	"testing"
)

/**
授权接口
*/

func TestSsoUrl(t *testing.T) {
	req := request.OpenSSOPrivilegeUrlRequest{}
	company := model.Company{}
	company.Name = "测试11-8-1"
	req.Company = &company
	user := model.User{}
	user.Contact = "10000000011"
	user.ContactType = "MOBILE"
	req.User = &user

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestAuthorizeUser(t *testing.T) {
	req := request.PersonalSignAuthUrlRequest{}
	user := model.User{}
	user.Contact = "10000000011"
	user.ContactType = "MOBILE"
	req.User = &user
	req.AuthDeadline = "2023-08-08"
	req.Remark = "想授权呀就授权"
	authInfo := model.PersonalSignUserAuthInfo{}
	authInfo.Mode = "IVS"
	authInfo.PaperType = "IDCARD"
	authInfo.Username = "宋一"
	authInfo.IdCardNo = "1231232131232131"
	authInfo.ModifyFields = []string{"IDCARDNO"}
	req.AuthInfo = &authInfo

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestDeauthorizeUser(t *testing.T) {
	req := request.PersonalSignDeauthRequest{}
	user := model.User{}
	user.Contact = "10000000011"
	user.ContactType = "MOBILE"
	req.User = &user

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestAuthorizeRecordQuery(t *testing.T) {
	req := request.PersonalSignAuthQueryRequest{}
	user := model.User{}
	user.Contact = "10000000011"
	user.ContactType = "MOBILE"
	req.User = &user

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}
