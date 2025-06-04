package model

type ImportForecastWeightPost struct {
	WaybillNumber string `json:"waybillNumber"`    // 运单号
	TotalWeight   string `json:"totalWeight"`      // 预报总重量(单位:g)
	Length        string `json:"length,omitempty"` // 包裹长(单位:cm)
	Width         string `json:"width,omitempty"`  // 包裹宽(单位:cm)
	Height        string `json:"height,omitempty"` // 包裹高(单位:cm)
}

type ImportForecastWeightResp struct {
	YWApiResponse
	Data any `json:"data"`
}
