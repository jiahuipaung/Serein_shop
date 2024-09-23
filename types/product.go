package types

type ProductListReq struct {
	CategoryID uint `form:"category_id" json:"category_id"`
	BasePage
}

type ProductResp struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	CategoryID    uint   `json:"category_id"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Price         string `json:"price"`
	// DiscountPrice string `json:"discount_price"`
	// View          uint64 `json:"view"`
	// CreatedAt     int64  `json:"created_at"`
	// Num           int    `json:"num"`
	// OnSale        bool   `json:"on_sale"`
}

type ProductCreateReq struct {
	ID  uint   `form:"id" json:"id"`
	Name          string `form:"name" json:"name"`
	CategoryID    uint   `form:"category_id" json:"category_id"`
	Title         string `form:"title" json:"title" `
	Info          string `form:"info" json:"info" `
	ImgPath       string `form:"img_path" json:"img_path"`
	Price         string `form:"price" json:"price"`
}

type ListProductImgReq struct {
	ID uint `json:"id" form:"id"`
}

type ProductShowReq struct {
	ID uint `form:"id" json:"id"`
}

type ProductSearchReq struct {
	ID uint `form:"id" json:"id"`
	Name          string `form:"name" json:"name"`
	CategoryID    int    `form:"category_id" json:"category_id"`
	Title         string `form:"title" json:"title" `
	Info          string `form:"info" json:"info" `
	Price         string `form:"price" json:"price"`
	// DiscountPrice string `form:"discount_price" json:"discount_price"`
	// OnSale        bool   `form:"on_sale" json:"on_sale"`
	BasePage
}

type ProductDeleteReq struct {
	ID uint `form:"id" json:"id"`
	BasePage
}
