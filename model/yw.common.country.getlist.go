package model

type YWCountryData struct {
	ID     string `json:"id"`     // 国家Id
	Code   string `json:"code"`   // 国家代码
	NameCh string `json:"nameCh"` // 国家中文名称
	NameEn string `json:"nameEn"` // 国家英文名称
}

type YWCountryResp struct {
	YWApiResponse
	Data []YWCountryData `json:"data"`
}
