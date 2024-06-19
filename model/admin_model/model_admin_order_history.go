package admin_model

import "time"

type ResOrderHistory struct {
	UserID         string    `json:"user_id,omitempty" db:"user_id"`
	Name           string    `json:"name,omitempty" db:"name"`
	Email          string    `json:"email,omitempty" db:"email"`
	PhoneNumber    string    `json:"phone_number,omitempty" db:"phone_number"`
	OrderID        int64     `json:"order_id,omitempty" db:"order_id"`
	OrderDate      time.Time `json:"order_date,omitempty" db:"order_date"`
	TotalPrice     float64   `json:"total_price,omitempty" db:"total_price"`
	Status         string    `json:"status,omitempty" db:"status"`
	Address        string    `json:"address,omitempty" db:"address"`
	PaymentMethod  string    `json:"payment_method,omitempty" db:"payment_method"`
	DiscountCodeID string    `json:"discount_code_id,omitempty" db:"discount_code_id"`
	ItemID         []string  `json:"item_id,omitempty" db:"item_id"`
	Quantity       []int     `json:"quantity,omitempty" db:"quantity"`
	Price          []float64 `json:"price,omitempty" db:"price"`
	ItemName       []string  `json:"item_name,omitempty" db:"item_name"`
	Description    []string  `json:"description,omitempty" db:"description"`
	ImageUrl       []string  `json:"image_url,omitempty" db:"image_url"`
}
