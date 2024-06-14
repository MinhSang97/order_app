package model

import (
	"encoding/json"
	"log"
)

type MenuItemsModel struct {
	ItemID              string    `json:"item_id" db:"item_id"`
	ItemName            string    `json:"item_name" db:"item_name"`
	Description         string    `json:"description" db:"description"`
	Price               float64   `json:"price" db:"price"`
	ImageUrl            string    `json:"image_url" db:"image_url"`
	CustomizationOption []string  `json:"customization_option" db:"customization_option"`
	ExtraPrice          []float64 `json:"extra_price" db:"extra_price"`
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
