package test

import (
	"fmt"
	"qiyuesuo/sdk/model"
	"qiyuesuo/sdk/request"
	"testing"
)

/**
外部客户接口
*/

func TestCustomerAdd(t *testing.T) {
	req := request.CustomerAddRequest{}
	req.TenantType = "COMPANY"
	req.CompanyName = "测试11-07-1"
	user := model.User{}
	user.Name = "宋一"
	user.Contact = "10000000281"
	user.ContactType = "MOBILE"
	req.User = &user

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}
