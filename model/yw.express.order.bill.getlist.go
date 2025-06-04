package model

type YWBilledOrderDetailPost struct {
	WaybillNumbers []string `json:"waybillNumbers"` // 运单号
}

type ShippingInfo struct {
	WaybillNumber  string  `json:"waybillNumber"`  // 运单号
	CustomerCode   string  `json:"customerCode"`   // 发货账号
	CalcWeight     int     `json:"calcWeight"`     // 计费重
	WeightUnit     string  `json:"weightUnit"`     // 重量单位 g
	ExpressLength  int     `json:"expressLength"`  // 包裹长（cm）
	ExpressWidth   int     `json:"expressWidth"`   // 包裹宽（cm）
	ExpressHeight  int     `json:"expressHeight"`  // 包裹高（cm）
	Currency       string  `json:"currency"`       // 币种 （CNY,USD,EUR 等）
	ArOriPrice     float64 `json:"arOriPrice"`     // 本币金额
	ArOriPriceToFc float64 `json:"arOriPriceToFc"` // 外币金额
	TransType      string  `json:"transType"`      // 交易类型
	TransTypeName  string  `json:"transTypeName"`  // 交易类型名称
}

type YWBilledOrderDetailResp struct {
	YWApiResponse
	Data []ShippingInfo `json:"data"`
}
