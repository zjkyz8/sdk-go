package test

import (
	"encoding/json"
	"fmt"
	"qiyuesuo/sdk/http"
	"qiyuesuo/sdk/model"
	"qiyuesuo/sdk/request"
	"testing"
)

var client = http.NewSaaSSdkClient("https://openapi.qiyuesuo.cn", "OcczsY1111", "okNDNASSo2eDAI8dZv0TF2DE111111")

func TestSaasCompanyDetail(t *testing.T) {
	req := request.SaaSCompanyDetailRequest{}
	req.CompanyName = "go测试公司-1"

	response, err := client.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestSaasCompanyAuthUrl(t *testing.T) {
	req := request.SaasCompanyAuthPageUrlRequest{}
	req.CompanyName = "go测试公司-1"
	applicantInfo := model.User{"10000000011", "MOBILE", "宋一"}
	applicantBs, _ := json.Marshal(applicantInfo)
	req.ApplicantInfo = string(applicantBs)

	response, err := client.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestSaasCompanyAuthTicket(t *testing.T) {
	req := request.SaaSCompanyAuthMiniappTicketRequest{}
	req.CompanyName = "go测试公司-2"
	applicantInfo := model.User{"10000000011", "MOBILE", "宋一"}
	applicantBs, _ := json.Marshal(applicantInfo)
	req.ApplicantInfo = string(applicantBs)

	response, err := client.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestSaasCompanyAuthResult(t *testing.T) {
	req := request.SaaSCompanyAuthResultRequest{}
	req.CompanyName = "go测试公司-1"

	response, err := client.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestSaasUserAuthUrl(t *testing.T) {
	req := request.SaaSUserAuthPageRequest{}
	req.Mode = "IVS"
	req.Username = "宋一"
	user := model.User{"10000000011", "MOBILE", ""}
	req.User = &user
	req.ModifyFields = []string{"USERNAME"}

	response, err := client.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestSaasUserAuthResult(t *testing.T) {
	req := request.SaaSUserAuthResultRequest{}
	user := model.User{"10000000011", "MOBILE", ""}
	req.User = &user
	//req.AuthId = "b5ceef37-ec6c-4efe-8cce-b1e1ac960338"

	response, err := client.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestSaasUserAuthTicket(t *testing.T) {
	req := request.SaaSUserAuthMiniappTicketRequest{}
	req.Mode = "IVS"
	req.Username = "宋一"
	user := model.User{"10000000011", "MOBILE", ""}
	req.User = &user
	req.ModifyFields = []string{"USERNAME"}

	response, err := client.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestSaasPrivilegeUrl(t *testing.T) {
	req := request.SaasPrivilegeUrlRequest{}
	req.CompanyId = "3123528615623541523"
	user := model.User{"10000000011", "MOBILE", ""}
	req.User = &user
	createToken := true
	req.CreateToken = &createToken
	req.CallbackUrl = "http://open.qiyuesuo.cn"

	response, err := client.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestSaasPrivilegeTicket(t *testing.T) {
	req := request.SaaSPrivilegeMiniappTicketRequest{}
	req.CompanyId = "3123528615623541523"
	user := model.User{"10000000011", "MOBILE", ""}
	req.User = &user

	response, err := client.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestSaasPrivilegeDetail(t *testing.T) {
	req := request.SaasPrivilegeDetailRequest{}
	req.CompanyId = "3123528615623541523"

	response, err := client.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestSaasApplyApp(t *testing.T) {
	req := request.SaasApplyAppRequest{}
	req.CompanyId = "3123528615623541523"
	req.AppName = "平台名称-go测试公司"
	req.Domain = "test.gosgf"
	req.LoginUrl = "https://test.gosgf/login"
	req.CasVerifyUrl = "https://test.gosgf/casverify"

	response, err := client.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestSaasResetToken(t *testing.T) {
	req := request.SaaSResetTokenRequest{}
	req.CompanyId = "3123528615623541523"

	response, err := client.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestSaaSPersonalAccessControl(t *testing.T) {
	req := request.SaaSPersonalAccessControlRequest{}
	user := model.User{"10000000011", "MOBILE", ""}
	req.User = &user
	req.Operate = "ENABLE"

	response, err := client.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestSaaSCompanyAccessControl(t *testing.T) {
	req := request.SaaSCompanyAccessControlRequest{}
	company := model.Company{Name: "go测试公司-1"}
	req.Company = &company
	req.Operate = "ENABLE"

	response, err := client.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestSaaSAppAccessControl(t *testing.T) {
	req := request.SaaSAppAccessControlRequest{}
	company := model.Company{Name: "go测试公司-1"}
	req.Company = &company
	req.Operate = "ENABLE"

	response, err := client.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestSaaSPersonalSignAuthorize(t *testing.T) {
	req := request.SaaSPersonalSignAuthUrlRequest{}
	user := model.User{"10000000011", "MOBILE", ""}
	req.User = &user
	req.AuthDeadline = "2023-10-02"
	req.Remark = "这是授权理由"
	company := model.Company{Name: "go测试公司-1"}
	req.Company = &company

	response, err := client.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestSaaSPersonalSignDeauthorize(t *testing.T) {
	req := request.SaaSPersonalSignDeauthRequest{}
	user := model.User{"10000000011", "MOBILE", ""}
	req.User = &user
	company := model.Company{Name: "go测试公司-1"}
	req.Company = &company

	response, err := client.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestSaaSCompanySignDeauthorize(t *testing.T) {
	req := request.SaaSSealSignAuthUrlRequest{}
	user := model.User{"10000000011", "MOBILE", ""}
	req.SealAdmin = &user
	company := model.Company{Name: "go测试公司-1"}
	req.Company = &company
	req.AuthDeadline = "2023-10-02"
	req.Remark = "这是授权理由"

	response, err := client.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}
