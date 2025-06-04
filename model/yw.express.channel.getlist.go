package model

type YWChannelPost struct {
	ChannelId string `json:"channelId"` // 产品编号（产品id）；默认不传查询所有交货仓；
}

type RespChannelData struct {
	NameCh string `json:"nameCh"`
	NameEn string `json:"nameEn"`
	Id     string `json:"id"`
}

// 响应
type YWChannelResp struct {
	YWApiResponse
	Data []RespChannelData `json:"data"`
}
