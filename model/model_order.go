package model

import (
	"encoding/json"
	"log"
	"time"
)

type OrderModel struct {
	OrderID        int64     `json:"order_id" db:"order_id"`
	OrderDate      time.Time `json:"order_date" db:"order_date"`
	TotalPrice     float64   `json:"total price" db:"total_price"`
	Status         string    `json:"status" db:"status"`
	Address        string    `json:"address" db:"address"`
	PaymentMethod  string    `json:"payment_method" db:"payment_method"`
	PaymentDate    time.Time `json:"payment_date" db:"payment_date"`
	Amount         float64   `json:"amount" db:"amount"`
	DiscountCodeId string    `json:"discount_code_id" db:"discount_code_id"`
	ItemID         []string  `json:"item_id" db:"item_id"`
	Quantity       []int     `json:"quantity" db:"quantity"`
	Price          []float64 `json:"price" db:"price"`
}

func (c *OrderModel) TableName() string {
	return "orders"
}

func (c *OrderModel) ToJson() string {
	bs, err := json.Marshal(c)
	if err != nil {
		log.Fatalln(err)

	}
	return string(bs)
}
func (c *OrderModel) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
}
