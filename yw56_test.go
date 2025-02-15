package yw56

import (
	"fmt"
	"github.com/zengweigg/yw56/config"
	"github.com/zengweigg/yw56/model"
	"testing"
)

func Test001(m *testing.T) {
	// 初始化
	cfg := config.LoadConfig()
	ywClient := NewYWService(*cfg)
	//构造测试请求数据
	postData := model.YWBlank{}
	resp, err := ywClient.Services.XiaoBao.QueryCountryList(postData)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("666", resp.YWApiResponse, resp.Code, resp.Data, resp.Msg, resp.Success)
}
