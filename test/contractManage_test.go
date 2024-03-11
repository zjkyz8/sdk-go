package test

import (
	"fmt"
	"os"
	"qiyuesuo/sdk/http"
	"qiyuesuo/sdk/model"
	"qiyuesuo/sdk/request"
	"testing"
)

/**
合同管理
*/

func TestDraft(t *testing.T) {
	contract := model.Contract{}
	contract.Subject = "go测试合同-08"
	user := model.User{}
	user.Name = "宋一"
	user.Contact = "10000000011"
	user.ContactType = "MOBILE"
	contract.Creator = &user
	category := model.Category{}
	category.Name = "平台-个人-预设流程"
	contract.Category = &category
	signatory1 := model.Signatory{}
	signatory1.TenantName = "测试11-8-1"
	signatory1.TenantType = "COMPANY"
	signatory1.Receiver = &user
	signatory1SerialNo := 1
	signatory1.SerialNo = &signatory1SerialNo
	signatory2 := model.Signatory{}
	signatory2.TenantName = "宋一"
	signatory2.TenantType = "PERSONAL"
	//delaySet := true
	//signatory2.DelaySet = &delaySet
	receiver := model.User{}
	receiver.Name = "宋一"
	receiver.Contact = "10000000011"
	receiver.ContactType = "MOBILE"
	signatory2.Receiver = &receiver
	signatory2SerialNo := 2
	signatory2.SerialNo = &signatory2SerialNo
	var signatories []*model.Signatory
	signatories = append(signatories, &signatory1)
	signatories = append(signatories, &signatory2)
	contract.Signatories = signatories
	var templateParams []*model.TemplateParam
	param1 := model.TemplateParam{}
	param1.Name = "param1"
	param1.Value = "v1"
	param2 := model.TemplateParam{}
	param2.Name = "param2"
	param2.Value = "v2"
	templateParams = append(templateParams, &param1)
	templateParams = append(templateParams, &param2)
	contract.TemplateParams = templateParams
	send := false
	contract.Send = &send

	req := request.ContractDraftRequest{}
	req.Contract = &contract
	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)

	result := response["result"].(map[string]interface{})
	fmt.Println("contractId:", result["id"])
}

func TestAddDocumentByFile(t *testing.T) {
	req := request.DocumentAddByFileRequest{}
	req.ContractId = "3121362688614077231"
	req.Title = "goFile1"
	req.FileSuffix = "pdf"
	sort := -1
	req.DocumentSort = &sort
	file, _ := os.Open("/Users/sgf/Downloads/文件模板三.pdf")
	req.File = &http.FileItem{file, ""}

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)

	result := response["result"].(map[string]interface{})
	fmt.Println("documentId:", result["documentId"])
}

func TestAddDocumentByMultiFile(t *testing.T) {
	req := request.DocumentAddByMultipleFileRequest{}
	req.ContractId = "3123924240072446838"
	req.Title = "go-multiFile"
	sort := -1
	req.DocumentSort = &sort
	file1, _ := os.Open("/Users/sgf/develop/临时/go/个人认证.pdf")
	file2, _ := os.Open("/Users/sgf/develop/临时/go/私有云错误码.pdf")
	var fis []*http.FileItem
	fis = append(fis, &http.FileItem{Src: file1})
	fis = append(fis, &http.FileItem{Src: file2})
	req.Files = fis

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
	result := response["result"].(map[string]interface{})
	fmt.Println("documentId:", result["documentId"])
}

func TestAddDocumentByTemplate(t *testing.T) {
	req := request.DocumentAddByTemplateRequest{}
	req.ContractId = "3119917770918068335"
	req.Title = "goTemplateFile5"
	req.TemplateId = "2766933842559242287"
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
	sort := -2
	req.DocumentSort = &sort

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
	result := response["result"].(map[string]interface{})
	fmt.Println("documentId:", result["documentId"])
}

func TestAddDocumentWatermark(t *testing.T) {
	req := request.ContractAddwartermarkRequest{}
	req.ContractId = "3121362688614077231"
	watermark1 := model.Watermark{}
	watermark1.Type_ = "TEXT"
	watermark1.Content = "gosdk测试水印22"
	var watermarks []*model.Watermark
	watermarks = append(watermarks, &watermark1)
	req.Watermarks = watermarks

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestSend(t *testing.T) {
	req := request.ContractSendRequest{}
	req.ContractId = "3121737020200780306"
	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestCopySend(t *testing.T) {
	req := request.ContractCopySendRequest{}
	req.ContractId = "3119917770918068335"
	copySendReceiver1 := model.CopySendReceiver{}
	copySendReceiver1.Name = "宋一"
	receiver1 := model.User{Contact: "10000000012", ContactType: "MOBILE"}
	copySendReceiver1.Receiver = &receiver1
	var copySendReceivers []*model.CopySendReceiver
	copySendReceivers = append(copySendReceivers, &copySendReceiver1)
	req.Receivers = copySendReceivers
	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestSignatoryEdit(t *testing.T) {
	req := request.SignatoryEditRequest{}
	req.SignatoryId = "3119917777083695222"
	req.TenantName = "宋一"
	user := model.User{}
	user.Contact = "10036350001"
	user.ContactType = "MOBILE"
	req.Receiver = &user
	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestSignatoryAdd(t *testing.T) {
	req := request.SignatoryAddRequest{}
	req.ContractId = "3119917770918068335"
	signatory := model.Signatory{}
	signatory.TenantName = "宋三"
	signatory.TenantType = "PERSONAL"
	receiver := model.User{}
	receiver.Name = "宋三"
	receiver.Contact = "10036350002"
	receiver.ContactType = "MOBILE"
	signatory.Receiver = &receiver
	signatorySerialNo := 3
	signatory.SerialNo = &signatorySerialNo
	req.Signatory = &signatory
	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestContractForceEnd(t *testing.T) {
	req := request.ContractForceEndRequest{}
	req.ContractId = "3119917770918068335"
	req.Reason = "强制结束的原因：没原因"
	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestContractForceEndAttachment(t *testing.T) {
	req := request.AddEndSignAttachmentRequest{}
	req.ContractId = "3119917770918068335"
	req.Title = "强制结束的附件"
	req.FileSuffix = "pdf"
	file, err := os.Open("/Users/sgf/Downloads/文件模板三.pdf")
	if err != nil {
		fmt.Println("读文件出错，", err.Error())
	} else {
		req.File = &http.FileItem{Src: file}
	}
	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestContractViewPage(t *testing.T) {
	req := request.ContractViewPageRequest{}
	req.ContractId = "3119917770918068335"
	//req.Lang = "en"
	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestContractList(t *testing.T) {
	req := request.ContractListRequest{}
	limit := 10
	req.SelectLimit = &limit
	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
	result := response["result"].(map[string]interface{})
	fmt.Println("totalCount:", result["totalCount"])
	list := result["list"].([]interface{})
	fmt.Println("list.len:", len(list))
}

func TestContractDelay(t *testing.T) {
	req := request.ContractDelayRequest{}
	req.ContractId = "3121362688614077231"
	days := 10
	req.Days = &days
	//req.ExpireDate = "2023-09-25"

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestContractStream(t *testing.T) {
	req := request.ContractStreamRequest{}
	req.ContractId = "3121362688614077231"

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
	result := response["result"]
	fmt.Println("totalCount:", result.(map[string]interface{})["totalCount"])
}

func TestContractDownloadUrl(t *testing.T) {
	req := request.ContractDownloadUrlRequest{}
	req.ContractId = "3121362688614077231"

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestContractDownload(t *testing.T) {
	req := request.ContractDownloadRequest{}
	req.ContractId = "3119917770918068335"
	req.DownloadItems = []string{"CONTRACT", "ATTACHMENT"}

	file, _ := os.OpenFile("/Users/sgf/develop/临时/go/goContract.zip", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	defer file.Close()
	err := sdkClient.Download(req, file)
	if err != nil {
		fmt.Println("error,", err.Error())
	}
}

func TestDocumentDownload(t *testing.T) {
	req := request.DocumentDownloadRequest{}
	req.DocumentId = "3120931077053550709"

	file, _ := os.OpenFile("/Users/sgf/develop/临时/go/go-doc1.pdf", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	defer file.Close()
	err := sdkClient.Download(req, file)
	if err != nil {
		fmt.Println("error,", err.Error())
	}
}

func TestDetail(t *testing.T) {
	req := request.ContractDetailRequest{}
	req.ContractId = "3122440559256474401"
	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
	result := response["result"].(map[string]interface{})
	fmt.Println("status:", result["status"])
	sigs := result["signatories"].([]interface{})
	fmt.Println(sigs[0].(map[string]interface{})["id"], sigs[0].(map[string]interface{})["tenantName"])
	fmt.Println(sigs[1].(map[string]interface{})["id"], sigs[1].(map[string]interface{})["tenantName"])
	docs := result["documents"].([]interface{})
	fmt.Println("docs:", docs)
	fmt.Println("docs.size:", len(docs))
	fmt.Println("doc1:", docs[0].(map[string]interface{})["id"])
}
