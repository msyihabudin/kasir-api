package models

// ProductWithCategory is used to return product with category name
type ProductWithCategory struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	Stock        int     `json:"stock"`
	CategoryName *string `json:"category_name"`
}
