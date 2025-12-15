package yw56

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/zengweigg/yw56/model"
)

type xiaoBaoService service

// 查询通达国家列表
func (s xiaoBaoService) QueryCountryList() (model.YWCountryResp, error) {
	postData := make(map[string]interface{})
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
	err = sonic.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return respData, nil
	}
	return respData, nil
}

// 创建运单
func (s xiaoBaoService) CreateExpressOrder(postData model.YWOrderPost) (model.YWOrderResp, error) {
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

// 查询交货仓列表
func (s xiaoBaoService) QueryWarehouseList(postData model.YWwarehousePost) (model.YWwarehouseResp, error) {
	respData := model.YWwarehouseResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("common.warehouse.getlist")
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

// 查询已开通的产品列表
func (s xiaoBaoService) QueryChannelList() (model.YWChannelResp, error) {
	respData := model.YWChannelResp{}
	// 请求数据
	postData := make(map[string]interface{})
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("express.channel.getlist")
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

// 打印标签
func (s xiaoBaoService) PrintLabel(postData model.YWPrintLabelPost) (model.YWPrintLabelResp, error) {
	respData := model.YWPrintLabelResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("express.order.label.get")
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

// 取消运单
func (s xiaoBaoService) CancelExpressOrder(postData model.YWCancelOrderPost) (model.YWCancelOrderPostOrderResp, error) {
	respData := model.YWCancelOrderPostOrderResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("express.order.cancel")
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

// 查询运单详情
func (s xiaoBaoService) QueryExpressOrderDetail(postData model.YWOrderDetailPost) (model.YWOrderDetailResp, error) {
	respData := model.YWOrderDetailResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("express.order.get")
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

// 批量查询单号详情
func (s xiaoBaoService) QueryExpressOrderDetailBatch(postData model.YWOrderDetailBatchPost) (model.YWOrderDetailBatchResp, error) {
	respData := model.YWOrderDetailBatchResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("express.order.getlist")
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

// 韩国个人通关码校验
func (s abroadService) CheckKoreaPersonalPassport(postData model.KoreaPersonalPassportPost) (model.KoreaPersonalPassportResp, error) {
	respData := model.KoreaPersonalPassportResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("common.verify.kr.pccc")
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

// 美国地址校验
func (s abroadService) VerifyUsAddress(postData model.UsAddressVerifyPost) (model.UsAddressVerifyResp, error) {
	respData := model.UsAddressVerifyResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("common.verify.us.address")
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

// 预报重量导入
func (s abroadService) ImportForecastWeight(postData model.ImportForecastWeightPost) (model.ImportForecastWeightResp, error) {
	respData := model.ImportForecastWeightResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("express.order.forecast_import")
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

// 已出账单详情
func (s abroadService) GetBilledOrderDetail(postData model.YWBilledOrderDetailPost) (model.YWBilledOrderDetailResp, error) {
	respData := model.YWBilledOrderDetailResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("express.order.bill.getlist")
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

// TEMU头程尾程面单上传
func (s abroadService) TemuCustomerLabelImport(postData model.YWTemuCustomerLabelPost) (model.YWTemuCustomerLabelResp, error) {
	respData := model.YWTemuCustomerLabelResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("express.order.customer_label.import")
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
