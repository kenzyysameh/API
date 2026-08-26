package models

type Order struct {
	ID          int         `json:"id" binding:"required"`
	UserID      int         `json:"user" binding:"required"`
	Items       []OrderItem `json:"items" binding:"required"`
	Total_Price float64     `json:"total_price" binding:"required"`
	Status      string      `json:"status"`
}
type OrderItem struct {
	ProductID int `json:"product_id" binding:"required"`
	Quantity  int `json:"quantity" binding:"required"`
}
