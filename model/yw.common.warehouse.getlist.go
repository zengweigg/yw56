package model

type YWwarehousePost struct {
	ChannelId string `json:"channelId,omitempty"` // 产品编号（产品id）；默认不传查询所有交货仓；
}

type RespWarehouseData struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Area string `json:"area"`
}

// 响应
type YWwarehouseResp struct {
	YWApiResponse
	Data []RespWarehouseData `json:"data"`
}
