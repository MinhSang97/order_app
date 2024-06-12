package payload

import (
	"encoding/json"
	"github.com/MinhSang97/order_app/model"
	"log"
	"time"
)

type DiscountCodesPayload struct {
	DiscountCodeID     string    `json:"discount_codes_id" validate:"required"`
	Title              string    `json:"title" validate:"required"`
	Description        string    `json:"description" validate:"required"`
	Code               string    `json:"code" validate:"required"`
	DiscountPercentage float64   `json:"discount_percentage" validate:"required"`
	ValidFrom          time.Time `json:"valid_from" validate:"required"`
	ValidTo            time.Time `json:"valid_to" validate:"required"`
	PromotionID        string    `json:"promotion_id" validate:"required"`
}

func (c *DiscountCodesPayload) ToModel() *model.DiscountCodesModel {
	discountCodesPayload := &model.DiscountCodesModel{
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

func (c *DiscountCodesPayload) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
}
