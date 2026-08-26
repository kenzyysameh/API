package models

type Review struct {
	ID           int    `json:"id"`
	ProductID    int    `json:"productid"`
	UserID       int    `json:"userid"`
	Rating       int    `json:"rating"`
	Comment      string `json:"comment"`
	CreationDate string `json:"creation_date"`
}
type NewReview struct {
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
}
