package test

import (
	"fmt"
	"qiyuesuo/sdk/request"
	"testing"
)

/**
业务分类接口
*/

func TestCategoryList(t *testing.T) {
	req := request.CategoryListRequest{}
	req.ModifyTimeStart = "2022-09-10 00:00:00"

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestCategoryDetail(t *testing.T) {
	req := request.CategoryDetailRequest{}
	req.CategoryId = "2594381609130005418"

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}
