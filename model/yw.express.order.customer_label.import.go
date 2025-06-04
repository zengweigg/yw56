package model

type YWTemuCustomerLabelPost struct {
	WaybillNumber  string `json:"waybillNumber"`         // 是 运单号
	TrackingNumber string `json:"trackingNumber"`        // 是 尾程单号
	LabelNumber    string `json:"labelNumber,omitempty"` // 否 标签单号
	LabelBase64    string `json:"labelBase64"`           // 是 面单标签PDF的Base64
}

// 响应
type YWTemuCustomerLabelResp struct {
	YWApiResponse
	Data any `json:"data"`
}
