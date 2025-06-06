package model

import "gopkg.in/guregu/null.v4"

// 收件人信息
type ReceiverInfo struct {
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Company     string `json:"company"`
	Country     string `json:"country"`
	State       string `json:"state"`
	City        string `json:"city"`
	ZipCode     string `json:"zipCode"`
	HouseNumber string `json:"houseNumber"`
	Address     string `json:"address"`
	TaxNumber   string `json:"taxNumber"`
}

// 发件人信息
type SenderInfo struct {
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Company     string `json:"company"`
	Country     string `json:"country"`
	State       string `json:"state"`
	City        string `json:"city"`
	ZipCode     string `json:"zipCode"`
	HouseNumber string `json:"houseNumber"`
	Address     string `json:"address"`
	TaxNumber   string `json:"taxNumber"`
}

type YWCsp struct {
	Csp string `json:"csp"`
}

// 进口清关信息
type ImportCustomsInfo struct {
	TaxPolicyExtends YWCsp `json:"taxPolicyExtends"`
}

// 创建运单
type YWOrderPost struct {
	ChannelID         string            `json:"channelId"`
	OrderSource       null.String       `json:"orderSource"`
	UserID            string            `json:"userId"`
	OrderNumber       string            `json:"orderNumber"`
	DateOfReceipt     string            `json:"dateOfReceipt"`
	Remark            string            `json:"remark"`
	ReceiverInfo      ReceiverInfo      `json:"receiverInfo"`      // 收件人信息
	ParcelInfo        ParcelInfo        `json:"parcelInfo"`        // 包裹信息
	SenderInfo        SenderInfo        `json:"senderInfo"`        // 发件人信息
	PoPStation        YWPointID         `json:"poPStation"`        // 自提网点信息
	ImportCustomsInfo ImportCustomsInfo `json:"importCustomsInfo"` // 进口清关信息
}

// 响应
type YWOrderResp struct {
	YWApiResponse
	Data struct {
		WaybillNumber     string `json:"waybillNumber"`
		OrderNumber       string `json:"orderNumber"`
		ReferenceNumber   string `json:"referenceNumber"`
		YanwenOrderNumber string `json:"yanwenOrderNumber"`
	} `json:"data"`
}
