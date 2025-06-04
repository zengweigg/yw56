package model

type YWPrintLabelPost struct {
	WaybillNumber string `json:"waybillNumber"`         // 运单号；
	PrintRemark   int    `json:"printRemark,omitempty"` // 是否单独打印拣货单 1:是 0:否(不传默认为否) 非必填；
}

// 响应
type YWPrintLabelResp struct {
	YWApiResponse
	Data struct {
		WaybillNumber string `json:"waybillNumber"` // 运单号
		ErrorMsg      string `json:"errorMsg"`      // 错误信息
		IsSuccess     bool   `json:"isSuccess"`     // 是否成功
		Base64String  string `json:"base64String"`  // 文件base64字符
	} `json:"data"`
}
