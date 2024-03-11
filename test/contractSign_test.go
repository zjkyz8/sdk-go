package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/zjkyz8/sdk-go/http"
	"github.com/zjkyz8/sdk-go/model"
	"github.com/zjkyz8/sdk-go/request"
)

/**
合同签署
*/

func TestCompanySign(t *testing.T) {
	req := request.ContractSignCompanyRequest{}
	param := model.SignParam{}
	param.ContractId = "3121362688614077231"
	stamper := model.Stamper{}
	stamper.Type_ = "COMPANY"
	stamper.DocumentId = "3121365707523162656"
	page := 1
	stamper.Page = &page
	offsetX := 0.1
	stamper.OffsetX = &offsetX
	offsetY := 0.1
	stamper.OffsetY = &offsetY
	var stampers []*model.Stamper
	stampers = append(stampers, &stamper)
	param.Stampers = stampers
	req.Param = &param

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestLpSign(t *testing.T) {
	req := request.ContractSignLpRequest{}
	param := model.SignParam{}
	param.ContractId = "3121742127780930396"
	stamper := model.Stamper{}
	stamper.Type_ = "LP"
	stamper.DocumentId = "3121742127378277206"
	page := 1
	stamper.Page = &page
	offsetX := 0.1
	stamper.OffsetX = &offsetX
	offsetY := 0.1
	stamper.OffsetY = &offsetY
	var stampers []*model.Stamper
	stampers = append(stampers, &stamper)
	param.Stampers = stampers
	req.Param = &param

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestAudit(t *testing.T) {
	req := request.ContractAuditRequest{}
	req.ContractId = "3121742127780930396"
	pass := true
	req.Pass = &pass
	req.Comment = "掐指一算，今日不宜通过"

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestNotice(t *testing.T) {
	req := request.ContractNoticeRequest{}
	req.ContractId = "3121656555125154667"

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestPageUrl(t *testing.T) {
	req := request.ContractPageRequest{}
	req.ContractId = "3121656555125154667"
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

func TestAppointPage(t *testing.T) {
	req := request.ContractAppointurlRequest{}
	req.ContractId = "3121658162923512178"

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestShortUrl(t *testing.T) {
	req := request.ContractShorturlRequest{}
	req.ContractId = "3121656555125154667"
	req.Contact = "10000000011"

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestInvalid(t *testing.T) {
	req := request.ContractInvalidRequest{}
	req.ContractId = "3122011527419339721"
	req.Reason = "作废原因"
	autoSign := false
	req.AutoSign = &autoSign
	req.SealId = "2730780295417565210"

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestInvalidByFile(t *testing.T) {
	req := request.ContractInvalidByFileRequest{}
	invalidInfo := model.InvalidInfo{}
	invalidInfo.ContractId = "3121737020200780306"
	invalidInfo.Title = "自定义作废文件"
	invalidInfo.FileSuffix = "pdf"
	invalidInfo.InvalidReason = "作废原因"
	stamper := model.Stamper{}
	stamper.SignatoryId = "3121737026265743894"
	stamper.Type_ = "COMPANY"
	page := 1
	stamper.Page = &page
	offsetX := 0.1
	stamper.OffsetX = &offsetX
	offsetY := 0.1
	stamper.OffsetY = &offsetY
	stamper2 := model.Stamper{}
	stamper2.SignatoryId = "3121737026416738842"
	stamper2.Type_ = "PERSONAL"
	page2 := 1
	stamper2.Page = &page2
	offsetX2 := 0.2
	stamper2.OffsetX = &offsetX2
	offsetY2 := 0.1
	stamper2.OffsetY = &offsetY2
	var stampers []*model.Stamper
	stampers = append(stampers, &stamper)
	stampers = append(stampers, &stamper2)
	invalidInfo.Stampers = stampers
	req.InvalidInfo = &invalidInfo
	file, _ := os.Open("/Users/sgf/Downloads/文件模板三.pdf")
	req.File = &http.FileItem{Src: file}
	user := model.User{}
	user.Name = "宋一"
	user.Contact = "10000000012"
	user.ContactType = "MOBILE"
	copySendReceiver := model.CopySendReceiver{}
	copySendReceiver.Name = "宋一"
	copySendReceiver.Receiver = &user
	var receivers []*model.CopySendReceiver
	receivers = append(receivers, &copySendReceiver)
	req.ReceiversInfo = receivers

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestInvalidByTemplate(t *testing.T) {
	req := request.InvalidByTemplateRequest{}
	req.ContractId = "3121742127780930396"
	req.TemplateId = "2766933842559242287"
	req.Title = "作废说明"
	// 指定模板参数
	var templateParams []*model.TemplateParam
	param1 := model.TemplateParam{}
	param1.Name = "param1"
	param1.Value = "v1"
	param2 := model.TemplateParam{}
	param2.Name = "param2"
	param2.Value = "v2"
	templateParams = append(templateParams, &param1)
	templateParams = append(templateParams, &param2)
	req.TemplateParams = templateParams
	// 指定签署位置
	stamper := model.Stamper{}
	stamper.SignatoryId = "3121742135292928864"
	stamper.Type_ = "COMPANY"
	page := 1
	stamper.Page = &page
	offsetX := 0.1
	stamper.OffsetX = &offsetX
	offsetY := 0.1
	stamper.OffsetY = &offsetY
	stamper2 := model.Stamper{}
	stamper2.SignatoryId = "3121742135469089636"
	stamper2.Type_ = "PERSONAL"
	page2 := 1
	stamper2.Page = &page2
	offsetX2 := 0.2
	stamper2.OffsetX = &offsetX2
	offsetY2 := 0.1
	stamper2.OffsetY = &offsetY2
	var stampers []*model.Stamper
	stampers = append(stampers, &stamper)
	stampers = append(stampers, &stamper2)
	req.Stampers = stampers

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestSignInvalid(t *testing.T) {
	req := request.ContractSignInvalidRequest{}
	req.ContractId = "3122011527419339721"
	req.SealId = "2730780295417565210"

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestSignUser(t *testing.T) {
	signParam := model.UserSignParam{}
	signParam.ContractId = "3122011527419339721"
	user := model.User{}
	user.Contact = "10000000011"
	user.ContactType = "MOBILE"
	signParam.User = &user
	req := request.ContractSignUserRequest{&signParam}

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestUserJoinurl(t *testing.T) {
	req := request.ContractPersonalJoinUrlRequest{}
	req.ContractId = "3122024039145611748"

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestRequiredParamList(t *testing.T) {
	req := request.RequiredParamListRequest{}
	req.ContractId = "3121651562611475189"
	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}
