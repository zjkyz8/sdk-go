package test

import (
	"fmt"
	"os"
	"qiyuesuo/sdk/model"
	"qiyuesuo/sdk/request"
	"testing"
)

/**
印章接口
*/

func TestSealList(t *testing.T) {
	req := request.SealListRequest{}
	req.ModifyTimeStart = "2021-05-25 00:00:00"

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestSealImage(t *testing.T) {
	req := request.SealImageRequest{}
	req.SealId = "2756508194640498724"

	file, _ := os.OpenFile("/Users/sgf/develop/临时/go/seal1.png", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	defer file.Close()
	err := sdkClient.Download(req, file)
	if err != nil {
		fmt.Println("error,", err.Error())
	}
}

func TestSealAutoCreate(t *testing.T) {
	req := request.SealAutoCreateRequest{}
	req.Name = "go自动生成印章2"
	sii := model.SealImageInfo{}
	sii.Style = "UNIVERSAL_SEAL"
	sii.Spec = "CIRCULAR_42"
	sii.Foot = "shallot"
	req.SealImageInfo = &sii
	user := model.User{}
	user.Contact = "10000000011"
	user.ContactType = "MOBILE"
	var users []*model.User
	users = append(users, &user)
	req.Users = users

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestSealDetail(t *testing.T) {
	req := request.SealDetailRequest{}
	req.SealId = "2756508194640498724"

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestSealEdit(t *testing.T) {
	req := request.SealEditRequest{}
	req.SealId = "3122087992198578926"
	req.SealName = "go测试印章1"

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestSealCreateByImage(t *testing.T) {
	req := request.SealCreatebyImageRequest{}
	req.SealName = "go图片印章"
	req.SealImage = "印章数据。。。"
	user := model.User{}
	user.Contact = "10000000011"
	user.ContactType = "MOBILE"
	var users []*model.User
	users = append(users, &user)
	req.SealUsers = users

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestSealStatusManage(t *testing.T) {
	req := request.SealStatusManageRequest{}
	req.SealId = "3122087992198578926"
	req.Operate = "DISABLE"

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestSealRemove(t *testing.T) {
	req := request.SealRemoveRequest{}
	req.SealId = "3122087992198578926"

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}
