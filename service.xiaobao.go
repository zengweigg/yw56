package yw56

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/zengweigg/yw56/model"
)

type xiaoBaoService service

// 查询通达国家列表
func (s xiaoBaoService) QueryCountryList(postData model.YWBlank) (model.YWCountryResp, error) {
	respData := model.YWCountryResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("common.country.getlist")
	fmt.Println(string(resp.Body()))
	if err != nil {
		return respData, err
	}
	// 解析数据
	err = jsoniter.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return respData, nil
	}
	return respData, nil
}
