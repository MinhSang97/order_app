package dto

import (
	"github.com/MinhSang97/order_app/payload"
)

type MenuItemsDto struct {
	ItemID              string    `json:"item_id" db:"item_id"`
	ItemName            string    `json:"item_name" db:"item_name"`
	Description         string    `json:"description" db:"description"`
	Price               float64   `json:"price" db:"price"`
	ImageUrl            string    `json:"image_url" db:"image_url"`
	CustomizationOption []string  `json:"customization_option"`
	ExtraPrice          []float64 `json:"extra_price"`
}

func (c *MenuItemsDto) ToPayLoad() *payload.MenuItemsPayload {
	menuItemsPayload := &payload.MenuItemsPayload{
		ItemID:              c.ItemID,
		ItemName:            c.ItemName,
		Description:         c.Description,
		Price:               c.Price,
		ImageUrl:            c.ImageUrl,
		CustomizationOption: c.CustomizationOption,
		ExtraPrice:          c.ExtraPrice,
	}
	return menuItemsPayload
}
