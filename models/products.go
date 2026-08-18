package models

type Product struct {
	Id          int     `json:"id" binding:"required"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"catergory"`
	Stock       int     `json:"stock"`
}

type ProductUpdateReq struct {
	Price *float64 `json:"price"`
	Stock *int     `json:"stock"`
}

var ProductsData = []Product{
	{
		Id:          1,
		Name:        "Wireless Headphones",
		Description: "Bluetooth headphones",
		Price:       1500,
		Category:    "electronics",
		Stock:       20,
	},
	{
		Id:          2,
		Name:        "Laptop",
		Description: "Lightweight laptop",
		Price:       25000,
		Category:    "electronics",
		Stock:       10,
	},
	{
		Id:          3,
		Name:        "Coffedebug Mug",
		Description: "Ceramic mug",
		Price:       200,
		Category:    "home",
		Stock:       50,
	},
}
