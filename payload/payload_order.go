package payload

import (
	"encoding/json"
	"github.com/MinhSang97/order_app/model"
	"log"
	"time"
)

type OrderPayload struct {
	OrderID        int64     `json:"order_id" db:"order_id"`
	OrderDate      time.Time `json:"order_date" db:"order_date"`
	TotalPrice     float64   `json:"total_price" db:"total_price"`
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

func (c *OrderPayload) ToModel() *model.OrderModel {
	orderPayload := &model.OrderModel{
		OrderID:        c.OrderID,
		OrderDate:      c.OrderDate,
		TotalPrice:     c.TotalPrice,
		Status:         c.Status,
		Address:        c.Address,
		PaymentMethod:  c.PaymentMethod,
		PaymentDate:    c.PaymentDate,
		Amount:         c.Amount,
		DiscountCodeId: c.DiscountCodeId,
		ItemID:         c.ItemID,
		Quantity:       c.Quantity,
		Price:          c.Price,
	}
	return orderPayload
}

func (c *OrderPayload) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
}
