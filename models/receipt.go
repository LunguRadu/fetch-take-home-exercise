package models

type Item struct {
	ShortDescription string `json:"shortDescription" binding:"required"`
	Price            string `json:"price" binding:"required"`
}

type Receipt struct {
	ID           string `json:"id"`
	UserId       int    `json:"userId" binding:"required"`
	Retailer     string `json:"retailer" binding:"required"`
	PurchaseDate string `json:"purchaseDate" binding:"required"`
	PurchaseTime string `json:"purchaseTime" binding:"required"`
	Items        []Item `json:"items" binding:"required"`
	Total        string `json:"total" binding:"required"`
	Points       int    `json:"points"`
}

type User struct {
	ID       int       `json:"id"`
	Name     string    `json:"name" binding:"required"`
	Receipts []Receipt `json:"receipts"`
	Points   int       `json:"points"`
}

type ReedemRequest struct {
	Points int `json:"points" binding:"required"`
}
