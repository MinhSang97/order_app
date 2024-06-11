package model

import (
	"encoding/json"
	"log"
)

type MenuItemsModel struct {
	ItemID                string  `json:"item_id" db:"item_id"`
	Name                  string  `json:"name" db:"name"`
	Description           string  `json:"description" db:"description"`
	Price                 float64 `json:"price" db:"price"`
	ImageUrl              string  `json:"image_url" db:"image_url"`
	CustomizationOption1  string  `json:"customization_option_1"`
	ExtraPrice1           float64 `json:"extra_price_1"`
	CustomizationOption2  string  `json:"customization_option_2"`
	ExtraPrice2           float64 `json:"extra_price_2"`
	CustomizationOption3  string  `json:"customization_option_3"`
	ExtraPrice3           float64 `json:"extra_price_3"`
	CustomizationOption4  string  `json:"customization_option_4"`
	ExtraPrice4           float64 `json:"extra_price_4"`
	CustomizationOption5  string  `json:"customization_option_5"`
	ExtraPrice5           float64 `json:"extra_price_5"`
	CustomizationOption6  string  `json:"customization_option_6"`
	ExtraPrice6           float64 `json:"extra_price_6"`
	CustomizationOption7  string  `json:"customization_option_7"`
	ExtraPrice7           float64 `json:"extra_price_7"`
	CustomizationOption8  string  `json:"customization_option_8"`
	ExtraPrice8           float64 `json:"extra_price_8"`
	CustomizationOption9  string  `json:"customization_option_9"`
	ExtraPrice9           float64 `json:"extra_price_9"`
	CustomizationOption10 string  `json:"customization_option_10"`
	ExtraPrice10          float64 `json:"extra_price_10"`
}

func (c *MenuItemsModel) TableName() string {
	return "menu_items"
}

func (c *MenuItemsModel) ToJson() string {
	bs, err := json.Marshal(c)
	if err != nil {
		log.Fatalln(err)

	}
	return string(bs)
}
func (c *MenuItemsModel) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
}
