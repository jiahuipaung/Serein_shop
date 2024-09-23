package types

type ListCategoryReq struct {
}

type ListCategoryResp struct {
	ID           uint   `json:"id"`
	CategoryName string `json:"category_name"`
	CreateAt     int64
}
