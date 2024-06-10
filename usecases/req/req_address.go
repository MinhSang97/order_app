package req

type ReqAddress struct {
	Address        string  `json:"address,omitempty" validate:"required"`
	Name           string  `json:"name,omitempty" validate:"required"`
	PhoneNumber    string  `json:"phone_number,omitempty" validate:"required"`
	TypeAddress    string  `json:"type_address,omitempty" validate:"required"`
	AddressDefault string  `json:"address_default,omitempty"`
	Lat            float64 `json:"lat,omitempty" validate:"required"`
	Long           float64 `json:"long,omitempty" validate:"required"`
	WardId         string  `json:"ward_id,omitempty" validate:"required"`
	WardText       string  `json:"ward_text,omitempty" validate:"required"`
	DistrictId     string  `json:"district_id,omitempty" validate:"required"`
	DistrictText   string  `json:"district_text,omitempty" validate:"required"`
	ProvinceId     string  `json:"province_id,omitempty" validate:"required"`
	ProvinceText   string  `json:"province_text,omitempty" validate:"required"`
	NationalId     string  `json:"national_id,omitempty" validate:"required"`
	NationalText   string  `json:"national_text,omitempty" validate:"required"`
}
