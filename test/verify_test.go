package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/zjkyz8/sdk-go/http"
	"github.com/zjkyz8/sdk-go/request"
)

/**
授权接口
*/

func TestVerify(t *testing.T) {
	req := request.VerifierRequest{}
	file, _ := os.Open("/Users/sgf/develop/临时/go/goContract/原文-goFile.pdf")
	defer file.Close()
	//s, err := io.ReadAll(file)
	//file2 := bytes.NewReader(s)
	req.File = &http.FileItem{Src: file}
	showImg := false
	req.ShowImg = &showImg

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}
