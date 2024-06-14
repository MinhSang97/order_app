package req

type ReqMenuItems struct {
	//ItemID      string  `json:"item_id" validate:"required"`
	Name                string    `json:"name" validate:"required"`
	Description         string    `json:"description" validate:"required"`
	Price               float64   `json:"price" validate:"required"`
	ImageUrl            string    `json:"image_url" validate:"required"`
	CustomizationOption []string  `json:"customization_option"`
	ExtraPrice          []float64 `json:"extra_price"`
}
