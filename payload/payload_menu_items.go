package payload

import (
	"encoding/json"
	"github.com/MinhSang97/order_app/model"
	"log"
)

type MenuItemsPayload struct {
	ItemID              string    `json:"item_id" db:"item_id"`
	Name                string    `json:"name" db:"name"`
	Description         string    `json:"description" db:"description"`
	Price               float64   `json:"price" db:"price"`
	ImageUrl            string    `json:"image_url" db:"image_url"`
	CustomizationOption []string  `json:"customization_option"`
	ExtraPrice          []float64 `json:"extra_price"`
}

func (c *MenuItemsPayload) ToModel() *model.MenuItemsModel {
	menuItemsPayload := &model.MenuItemsModel{
		ItemID:              c.ItemID,
		Name:                c.Name,
		Description:         c.Description,
		Price:               c.Price,
		ImageUrl:            c.ImageUrl,
		CustomizationOption: c.CustomizationOption,
		ExtraPrice:          c.ExtraPrice,
	}

	return menuItemsPayload
}

func (c *MenuItemsPayload) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
}
