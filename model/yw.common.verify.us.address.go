package model

type UsPassportData struct {
	Address string `json:"address"` // 收件人地址
	ZipCode string `json:"zipCode"` // 收件人邮编
	City    string `json:"city"`    // 收件人城市
	State   string `json:"state"`   // 收件人州（建议二字码）
}

type UsAddressVerifyPost struct {
	ReceiverInfo UsPassportData `json:"receiverInfo"`
}

type UsAddressVerify struct {
	ReceiverInfo struct {
		Address  string `json:"address"`
		City     string `json:"city"`
		State    string `json:"state"`
		ZipCode4 string `json:"zipCode4"`
		ZipCode5 string `json:"zipCode5"`
	} `json:"receiverInfo"`
}
type UsAddressVerifyResp struct {
	YWApiResponse
	Data UsAddressVerify `json:"data"`
}
