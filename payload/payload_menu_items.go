package payload

import (
	"encoding/json"
	"github.com/MinhSang97/order_app/model"
	"log"
)

type MenuItemsPayload struct {
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

func (c *MenuItemsPayload) ToModel() *model.MenuItemsModel {
	menuItemsPayload := &model.MenuItemsModel{
		ItemID:                c.ItemID,
		Name:                  c.Name,
		Description:           c.Description,
		Price:                 c.Price,
		ImageUrl:              c.ImageUrl,
		CustomizationOption1:  c.CustomizationOption1,
		ExtraPrice1:           c.ExtraPrice1,
		CustomizationOption2:  c.CustomizationOption2,
		ExtraPrice2:           c.ExtraPrice2,
		CustomizationOption3:  c.CustomizationOption3,
		ExtraPrice3:           c.ExtraPrice3,
		CustomizationOption4:  c.CustomizationOption4,
		ExtraPrice4:           c.ExtraPrice4,
		CustomizationOption5:  c.CustomizationOption5,
		ExtraPrice5:           c.ExtraPrice5,
		CustomizationOption6:  c.CustomizationOption6,
		ExtraPrice6:           c.ExtraPrice6,
		CustomizationOption7:  c.CustomizationOption7,
		ExtraPrice7:           c.ExtraPrice7,
		CustomizationOption8:  c.CustomizationOption8,
		ExtraPrice8:           c.ExtraPrice8,
		CustomizationOption9:  c.CustomizationOption9,
		ExtraPrice9:           c.ExtraPrice9,
		CustomizationOption10: c.CustomizationOption10,
		ExtraPrice10:          c.ExtraPrice10,
	}

	return menuItemsPayload
}

func (c *MenuItemsPayload) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
}
