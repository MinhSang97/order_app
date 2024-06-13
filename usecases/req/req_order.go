package req

import "time"

type ReqOrder struct {
	//OrderID        string    `json:"order_id" validate:"required"`
	OrderDate      time.Time `json:"order_date" validate:"required"`
	TotalPrice     float64   `json:"total_price" validate:"required"`
	Status         string    `json:"status" validate:"required"`
	Address        string    `json:"address" validate:"required"`
	PaymentMethod  string    `json:"payment_method" validate:"required"`
	PaymentDate    time.Time `json:"payment_date" validate:"required"`
	Amount         float64   `json:"amount" validate:"required"`
	DiscountCodeId string    `json:"discount_code_id" validate:"required"`
	ItemID         []string  `json:"item_id" validate:"required"`
	Quantity       []int     `json:"quantity" validate:"required"`
	Price          []float64 `json:"price" validate:"required"`
}

type OrderItem struct {
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}
