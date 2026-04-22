package dto

type ProductDTO struct {
	Name  string `json:"product_name" binding:"required"`
	Price int    `json:"product_price" binding:"required"`
	Stock int    `json:"product_stock" binding:"required"`
}

type productRequest struct {
	Id    int    `json:"product_id" binding:"required"`
	Name  string `json:"product_name"`
	Price int    `json:"product_price"`
	Stock int    `json:"product_stock"`
}
