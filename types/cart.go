package types

type CartServiceReq struct {
	Id       uint `form:"id" json:"id"`
	ProuctID uint `form:"product_id" json:"product_id"`
	Num      uint `form:"num" json:"num"`
	UserID   uint `form:"user_id" json:“user_id”`
}

type CartCreateReq struct {
	ProductID uint `form:"product_id" json:"product_id"`
}

type CartDeleteReq struct {
	Id uint `form:"id" json:"id"`
}

type UpdateCartServiceReq struct {
	ID  uint `form:"id" json:"id"`
	Num uint `form:"num" json:"num"`
}

type CartListReq struct {
	BasePage
}

// 购物车
type CartResp struct {
	ID            uint   `json:"id"`
	UserID        uint   `json:"user_id"`
	ProductID     uint   `json:"product_id"`
	CreatedAt     int64  `json:"created_at"`
	Num           uint   `json:"num"`
	MaxNum        uint   `json:"max_num"`
	Check         bool   `json:"check"`
	Name          string `json:"name"`
	ImgPath       string `json:"img_path"`
	DisCountPrice string `json:"discount_price"`
	Info          string `json:"info"`
}
