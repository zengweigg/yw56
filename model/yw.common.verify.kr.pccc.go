package model

// 韩国个人通关码校验

type KoreaPersonalPassportData struct {
	Name      string `json:"name"`      // 收件人姓名
	Phone     string `json:"phone"`     // 收件人联系方式
	TaxNumber string `json:"taxNumber"` // 收件人税号
}

type KoreaPersonalPassportPost struct {
	ReceiverInfo KoreaPersonalPassportData `json:"receiverInfo"`
}

// 响应
type KoreaPersonalPassportResp struct {
	YWApiResponse
	Data any `json:"data"`
}
