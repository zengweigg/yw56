package yw56

import (
	"fmt"
	"github.com/zengweigg/yw56/config"
	"github.com/zengweigg/yw56/model"
	"gopkg.in/guregu/null.v4"
	"testing"
)

func Test001(m *testing.T) {
	// 初始化
	cfg := config.LoadConfig()
	ywClient := NewYWService(*cfg)
	//  追踪物流
	apikey := cfg.APIKey
	res1, _ := ywClient.Services.Track.Tracking(apikey, "UH183870291YP")
	fmt.Println("响应", res1.Code, res1.Message, res1.Result, res1.RequestTime)

	// 测试
	res, _ := ywClient.Services.XiaoBao.QueryCountryList()
	fmt.Println(res)

	// 构造下单请求数据
	receiverInfo := model.ReceiverInfo{
		Name:        "glassware", // 收件人姓名填写不规范，仅支持英文，空格
		Phone:       "18231730589",
		Email:       "529932299@qq.com",
		Company:     "yanwen",
		Country:     "115",
		State:       "he bei sheng",
		City:        "cang zhou shi",
		ZipCode:     "10110",
		HouseNumber: "#4501124",
		Address:     "he fang jie cang zhou kai fa qu",
		TaxNumber:   "qwer123",
	}
	senderInfo := model.SenderInfo{
		Name:        "glasswarb",
		Phone:       "18231730588",
		Email:       "529932298@qq.com",
		Company:     "yanwen",
		Country:     "115",
		State:       "he bei sheng",
		City:        "cang zhou shi",
		ZipCode:     "101110",
		HouseNumber: "#4501124",
		Address:     "he fang jie cang zhou kai fa qu",
		TaxNumber:   "qwer123",
	}
	productList := make([]model.Product, 1)
	productList[0] = model.Product{
		GoodsNameCh: "杯子",
		GoodsNameEn: "CUP",
		Price:       "12.5",
		Hscode:      "8400000001",
		URL:         "http://www.aliexpress.com/item//32681820727.html",
		Material:    "material",
		Quantity:    "10",
		Weight:      "1000",
		Imei:        "IMEI1001",
	}
	parcelInfo := model.ParcelInfo{
		ProductList:   productList,
		HasBattery:    1,
		Currency:      "USD",
		TotalPrice:    "50.01",
		TotalQuantity: "10",
		TotalWeight:   "243",
		Height:        "10",
		Width:         "10",
		Length:        "10",
		Ioss:          "123456",
	}
	csp := model.YWCsp{
		Csp: "19237423SDFGEG",
	}
	importCustomsInfo := model.ImportCustomsInfo{
		TaxPolicyExtends: csp,
	}
	pointID := model.YWPointID{
		PointID: "1001",
	}
	postData := model.YWOrderPost{
		ChannelID:         "481",
		OrderSource:       null.StringFrom("portal"),
		UserID:            "99000014",
		OrderNumber:       "KI1000000001B",
		DateOfReceipt:     "2025-05-31",
		Remark:            "拣货单信息",
		ReceiverInfo:      receiverInfo,
		SenderInfo:        senderInfo,
		ParcelInfo:        parcelInfo,
		PoPStation:        pointID,
		ImportCustomsInfo: importCustomsInfo,
	}
	fmt.Println(666, postData)
	// resp, err := ywClient.Services.XiaoBao.CreateExpressOrder(postData)
	// if err != nil {
	// 	fmt.Printf("Error: %v\n", err)
	// 	return
	// }
	// fmt.Println("响应", resp.Code, resp.Msg, resp.Success)
	// fmt.Println("数据", resp.Data.WaybillNumber, resp.Data.OrderNumber, resp.Data.ReferenceNumber, resp.Data.YanwenOrderNumber)
}
