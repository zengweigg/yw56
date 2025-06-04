package model

type AbdOrderCancelPost struct {
	WaybillNumber string `json:"waybillNumber"` // 运单号
	Note          string `json:"note"`          // 取消原因
}

type AbdOrderCancelResp struct {
	YWApiResponse
	Data any `json:"data"`
}
