package dto

import (
	"github.com/MinhSang97/order_app/payload"
	"time"
)

type OrderDto struct {
	OrderID        int64          `json:"order_id" db:"order_id"`
	OrderDate      time.Time      `json:"order_date" db:"order_date"`
	TotalPrice     float64        `json:"total_price" db:"total_price"`
	Status         string         `json:"status" db:"status"`
	Address        string         `json:"address" db:"address"`
	PaymentMethod  string         `json:"payment_method" db:"payment_method"`
	PaymentDate    time.Time      `json:"payment_date" db:"payment_date"`
	Amount         float64        `json:"amount" db:"amount"`
	DiscountCodeId string         `json:"discount_code_id" db:"discount_code_id"`
	ItemID         []string       `json:"item_id" db:"item_id"`
	Quantity       []int          `json:"quantity" db:"quantity"`
	Price          []float64      `json:"price" db:"price"`
	Items          []OrderItemDto `json:"items"`
}
type OrderItemDto struct {
	ItemID   string  `json:"item_id"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

func (c *OrderDto) ToPayload() *payload.OrderPayload {
	orderDto := &payload.OrderPayload{
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
	return orderDto
}
