package yw56

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/zengweigg/yw56/model"
)

type abroadService service

// 创建运单
func (s abroadService) CreateExpressOrder(postData model.YWOrderPost) (model.YWOrderResp, error) {
	respData := model.YWOrderResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("express.order.create")
	fmt.Println(string(resp.Body()))
	if err != nil {
		return respData, err
	}
	// 解析数据
	err = sonic.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return respData, nil
	}
	return respData, nil
}
