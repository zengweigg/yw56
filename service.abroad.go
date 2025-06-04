package yw56

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/zengweigg/yw56/model"
)

type abroadService service

// 尾程业务-取消运单
func (s abroadService) AbdOrderCancel(postData model.AbdOrderCancelPost) (model.AbdOrderCancelResp, error) {
	respData := model.AbdOrderCancelResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("express.order.cancel")
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
