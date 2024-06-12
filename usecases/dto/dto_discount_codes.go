package dto

import (
	"github.com/MinhSang97/order_app/payload"
	"time"
)

type DiscountCodesDto struct {
	DiscountCodeID     string    `json:"discount_code_id" db:"discount_codes_id"`
	Title              string    `json:"title" db:"title"`
	Description        string    `json:"description" db:"description"`
	Code               string    `json:"code" db:"code"`
	DiscountPercentage float64   `json:"discount_percentage" db:"discount_percentage"`
	ValidFrom          time.Time `json:"valid_from" db:"valid_from"`
	ValidTo            time.Time `json:"valid_to" db:"valid_to"`
	PromotionID        string    `json:"promotion_id" db:"promotion_id"`
}

func (c *DiscountCodesDto) ToPayLoad() *payload.DiscountCodesPayload {
	discountCodesPayload := &payload.DiscountCodesPayload{
		DiscountCodeID:     c.DiscountCodeID,
		Title:              c.Title,
		Description:        c.Description,
		Code:               c.Code,
		DiscountPercentage: c.DiscountPercentage,
		ValidFrom:          c.ValidFrom,
		ValidTo:            c.ValidTo,
		PromotionID:        c.PromotionID,
	}
	return discountCodesPayload
}
