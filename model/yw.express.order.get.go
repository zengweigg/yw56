package model

type YWOrderDetailPost struct {
	WaybillNumber string `json:"waybillNumber"` // 运单号
}

// 收件人信息
type YwReceiverInfo struct {
	Name        string `json:"name"`
	Company     string `json:"company"`
	CountryID   string `json:"countryId"`
	CountryName string `json:"countryName"`
	Phone       string `json:"phone"`
	State       string `json:"state"`
	City        string `json:"city"`
	Email       string `json:"email"`
	ZipCode     string `json:"zipCode"`
	TaxNumber   string `json:"taxNumber"`
	Address     string `json:"address"`
	HouseNumber string `json:"houseNumber"`
}

// 发件人信息
type YwSenderInfo struct {
	Name        string `json:"name"`
	Company     string `json:"company"`
	CountryID   string `json:"countryId"`
	CountryName string `json:"countryName"`
	Phone       string `json:"phone"`
	State       string `json:"state"`
	City        string `json:"city"`
	Email       string `json:"email"`
	ZipCode     string `json:"zipCode"`
	TaxNumber   string `json:"taxNumber"`
	Address     string `json:"address"`
	HouseNumber string `json:"houseNumber"`
}

type TaxPolicyExtends struct {
	Csp string `json:"csp"` // 瑞士CSP deferment account ID
}

// 进口清关信息
type YwImportCustomsInfo struct {
	TaxPolicyExtends TaxPolicyExtends `json:"taxPolicyExtends"` // 国家相关税收政策信息
}

type Waybill struct {
	WaybillNumber     string              `json:"waybillNumber"`   // 运单号
	OrderNumber       string              `json:"orderNumber"`     // 订单号
	ReferenceNumber   string              `json:"referenceNumber"` // 转单号
	YanwenNumber      string              `json:"yanwenNumber"`
	YanwenOrderNumber string              `json:"yanwenOrderNumber"` // 燕文流水号(唯一)
	UserID            string              `json:"userId"`            // 发货账号
	OrderSource       string              `json:"orderSource"`       // 订单来源
	ChannelID         string              `json:"channelId"`         // 产品编号(产品id)
	ChannelName       string              `json:"channelName"`       // 产品名称
	DateOfReceipt     string              `json:"dateOfReceipt"`     // 收款到账日期
	CreateTime        string              `json:"createTime"`        // 下单时间
	CompanyCode       string              `json:"companyCode"`       // 发货账号所属燕文揽收仓编码
	IsPrint           int                 `json:"isPrint"`           // 是否打印 1:是 0:否
	Remark            string              `json:"remark"`            // 拣货单信息
	Status            int                 `json:"status"`            // 运单状态：0 已制单,1 已确认发货,2 已收货,3 运输中,4已妥投,5已取消,6 已截留,7 投递失败,8 仓内异常,  9 仓内退件,10 退件签收,11 转运异常,12 派送中,13 待提取,15 追踪结束；
	SalesPlatform     string              `json:"salesPlatform"`     // 销售平台
	HandoverCode      string              `json:"handoverCode"`      // 海外交货地
	ReceiverInfo      YwReceiverInfo      `json:"receiverInfo"`      // 收件人信息
	SenderInfo        YwSenderInfo        `json:"senderInfo"`        // 发件人信息
	ParcelInfo        ParcelInfo          `json:"parcelInfo"`        // 包裹信息
	PoPStation        YWPointID           `json:"poPStation"`        // 自提网点信息
	ImportCustomsInfo YwImportCustomsInfo `json:"importCustomsInfo"` // 进口清关信息
}

// 响应
type YWOrderDetailResp struct {
	YWApiResponse
	Data Waybill `json:"data"`
}
