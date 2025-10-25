package model

type YWApiResponse struct {
	Msg     string `json:"message"` // 消息内容
	Code    string `json:"code"`    // 消息编码 0:成功 >0:失败 <0:系统异常
	Success bool   `json:"success"` // 是否成功 true:成功 false:失败
}

// 包裹信息
type ParcelInfo struct {
	ProductList   []Product `json:"productList"`
	HasBattery    int       `json:"hasBattery"`    // 是否含电 1:是 0:否
	Currency      string    `json:"currency"`      // 申报币种
	TotalPrice    string    `json:"totalPrice"`    // 申报总价值：汇总多组商品信息品名申报单价*数量之和
	TotalQuantity int       `json:"totalQuantity"` // 申报总数量
	TotalWeight   int       `json:"totalWeight"`   // 总重量(单位:g)
	Height        int       `json:"height"`        // 包裹高度(单位:cm)
	Width         int       `json:"width"`         // 包裹宽度(单位:cm)
	Length        int       `json:"length"`        // 包裹长度(单位:cm)
	Ioss          string    `json:"ioss"`          // IOSS税号
}

type Product struct {
	GoodsNameCh string `json:"goodsNameCh"` // 中文品名
	GoodsNameEn string `json:"goodsNameEn"` // 英文品名
	Price       string `json:"price"`       // 申报单价
	Hscode      string `json:"hscode"`      // 商品海关编码
	URL         string `json:"url"`         // 商品链接
	Material    string `json:"material"`    // 商品材质
	Quantity    string `json:"quantity"`    // 申报数量
	Weight      int `json:"weight"`      // 单票重量(单位:g)
	Imei        string `json:"imei"`        // IMEI编码
}

// 自提网点信息
type YWPointID struct {
	PointID string `json:"pointId"` // 自提点ID
}
