package req

type ReqSignUp struct {
	Email        string  `json:"email,omitempty" validate:"required,email"`
	PassWord     string  `json:"password,omitempty" validate:"required"`
	Name         string  `json:"name,omitempty" validate:"required"`
	PhoneNumber  string  `json:"phone_number,omitempty" validate:"required"`
	Address      string  `json:"address,omitempty" validate:"required"`
	Telegram     string  `json:"telegram,omitempty" validate:"required"`
	Lat          float64 `json:"lat,omitempty"`
	Long         float64 `json:"long,omitempty"`
	WardId       string  `json:"ward_id,omitempty"`
	WardText     string  `json:"ward_text,omitempty"`
	DistrictId   string  `json:"district_id,omitempty"`
	DistrictText string  `json:"district_text,omitempty"`
	ProvinceId   string  `json:"province_id,omitempty"`
	ProvinceText string  `json:"province_text,omitempty"`
	NationalId   string  `json:"national_id,omitempty"`
	NationalText string  `json:"national_text,omitempty"`
}
