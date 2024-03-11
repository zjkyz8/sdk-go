package test

import (
	"fmt"
	"qiyuesuo/sdk/model"
	"qiyuesuo/sdk/request"
	"testing"
)

/**
个人认证接口
*/

func TestUserAuthPage(t *testing.T) {
	req := request.UserAuthPageRequest{}
	user := model.User{}
	user.Name = "宋一"
	user.Contact = "10000000281"
	user.ContactType = "MOBILE"
	req.User = &user
	req.Mode = "FACE"

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestUserAuthResult(t *testing.T) {
	req := request.UserAuthResultRequest{}
	req.Contact = "10000000281"
	req.ContactType = "MOBILE"

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}
