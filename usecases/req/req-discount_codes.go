package req

import (
	"time"
)

type ReqDiscountCodes struct {
	DiscountCodeID     string    `json:"discount_code_id" validate:"required"`
	Title              string    `json:"title" validate:"required"`
	Description        string    `json:"description" validate:"required"`
	Quantity           int       `json:"quantity" validate:"required"`
	Code               string    `json:"code" validate:"required"`
	DiscountPercentage float64   `json:"discount_percentage" validate:"required"`
	ValidFrom          time.Time `json:"valid_from" validate:"required"`
	ValidTo            time.Time `json:"valid_to" validate:"required"`
	PromotionID        string    `json:"promotion_id" validate:"required"`
}
