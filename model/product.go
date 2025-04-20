package model

type ProductInsertReq struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float32 `json:"price" binding:"required"`
	Qty         int     `json:"qty" binding:"required"`
}

type ProductInsertRes struct {
	Msg string `json:"msg"`
}
