package model

import (
	"encoding/json"
	"log"
	"time"
)

type DiscountCodesModel struct {
	DiscountCodeID     string    `json:"discount_code_id" db:"discount_codes_id"`
	Title              string    `json:"title" db:"title"`
	Description        string    `json:"description" db:"description"`
	Code               string    `json:"code" db:"code"`
	DiscountPercentage float64   `json:"discount_percentage" db:"discount_percentage"`
	ValidFrom          time.Time `json:"valid_from" db:"valid_from"`
	ValidTo            time.Time `json:"valid_to" db:"valid_to"`
	PromotionID        string    `json:"promotion_id" db:"promotion_id"`
}

func (c *DiscountCodesModel) TableName() string {
	return "discount_codes"
}

func (c *DiscountCodesModel) ToJson() string {
	bs, err := json.Marshal(c)
	if err != nil {
		log.Fatalln(err)

	}
	return string(bs)
}
func (c *DiscountCodesModel) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
}
