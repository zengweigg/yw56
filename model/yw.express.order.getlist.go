package model

type YWOrderDetailBatchPost struct {
	ListNumber []string `json:"listNumber"` // 查询单号集合，支持运单号或订单号，最多查询50个
}

type YWOrderDetailBatchResp struct {
	YWApiResponse
	Data []Waybill `json:"data"`
}
