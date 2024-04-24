package entities

type Product struct {
	ID       string  `json:"product_id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	SellerId string  `json:"seller_id"`
}
