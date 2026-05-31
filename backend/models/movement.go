package models

type movement struct {
	ID       uint   `json:"id"`
	Product  uint   `json:"product_id"`
	Type     string `json:"type"` // entrada ou saida
	Quantity int    `json:"quantity"`
}
