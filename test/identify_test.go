package test

import (
	"fmt"
	"testing"

	"github.com/zjkyz8/sdk-go/request"
)

func TestUserIdentifyVerify(t *testing.T) {
	req := request.IdentityVerifyRequest{}
	req.Name = "宋一"
	req.Mobile = "10000000011"
	req.CardNo = "410000000000000009"

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestCompanyIdentifyVerify(t *testing.T) {
	req := request.CompanyInfoVerifyRequest{}
	req.CompanyName = "上海亘岩网络科技有限公司"
	req.RegisterNo = "91310120MA1HKGA51W"
	req.LegalPerson = "衡晓辉"

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}
