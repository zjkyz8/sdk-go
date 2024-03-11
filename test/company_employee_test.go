package test

import (
	"fmt"
	"qiyuesuo/sdk/model"
	"qiyuesuo/sdk/request"
	"testing"
)

/**
组织架构接口
*/

func TestCompanyDetail(t *testing.T) {
	req := request.CompanyDetailRequest{}
	req.CompanyName = "测试11-8-1"
	//req.RegisterNo = "91310120MA1HKGA51W"

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestPlatformInfo(t *testing.T) {
	req := request.PlatformDetailRequest{}

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestListSubCompany(t *testing.T) {
	req := request.SubCompanyListRequest{}

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestCompanyChangePage(t *testing.T) {
	req := request.CompanyChangeInfoRequest{}
	req.CompanyName = "测试11-07-1"
	user := model.User{}
	user.Contact = "10000000011"
	user.ContactType = "MOBILE"
	req.Applicant = &user

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestListEmployee(t *testing.T) {
	req := request.EmployeeListRequest{}
	//req.TenantName = "测试11-8-1"
	selectLimit := 3
	req.SelectLimit = &selectLimit

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
	result := response["result"].(map[string]interface{})
	fmt.Println("查询的数量", len(result["list"].([]interface{})))
	fmt.Println("总数量", result["totalCount"])
}

func TestAddEmployee(t *testing.T) {
	req := request.EmployeeCreateRequest{}
	user := model.User{}
	user.Name = "宋三"
	user.Contact = "10000000281"
	user.ContactType = "MOBILE"
	req.User = &user
	req.Number = "天宫一号"

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestUpdateEmployee(t *testing.T) {
	req := request.EmployeeUpdateRequest{}
	user := model.User{}
	user.Contact = "10000000281"
	user.ContactType = "MOBILE"
	req.User = &user
	req.Number = "银河一号"

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestRemoveEmployee(t *testing.T) {
	req := request.EmployeeRemoveRequest{}
	user := model.User{}
	user.Contact = "10000000281"
	user.ContactType = "MOBILE"
	req.Param = &user

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestEmployeeDetail(t *testing.T) {
	req := request.EmployeeDetailRequest{}
	user := model.User{}
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

func TestInviteSumCompany(t *testing.T) {
	req := request.SubCompanyInviteRequest{}
	req.CompanyName = "测试11-07-1"

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestRemoveSumCompany(t *testing.T) {
	req := request.SubCompanyRemoveRequest{}
	req.CompanyName = "测试11-07-1"

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestManageRoles(t *testing.T) {
	req := request.RoleManagementRequest{}
	req.Role = "SEAL_ADMIN"
	req.Operate = "REMOVE"
	user := model.User{}
	user.Contact = "10000000281"
	user.ContactType = "MOBILE"
	users := []*model.User{}
	users = append(users, &user)
	req.Users = users

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}
