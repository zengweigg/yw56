package model

type YWCancelOrderPost struct {
	WaybillNumber string `json:"waybillNumber"`  // 运单号
	Note          string `json:"note,omitempty"` // 取消原因 可选
}

// 响应
type YWCancelOrderPostOrderResp struct {
	YWApiResponse
	Data any `json:"data"`
}
