package domain

type Order struct {
	ID       string  `json:"id"`
	Customer string  `json:"customer"`
	Total    float64 `json:"total"`
}
