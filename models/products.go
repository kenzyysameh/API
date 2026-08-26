package models

type Product struct {
	Id          int     `json:"id" `
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
