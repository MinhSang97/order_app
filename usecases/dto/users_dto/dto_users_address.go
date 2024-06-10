package users_dto

import "github.com/MinhSang97/order_app/payload/users_payload"

type UsersAddressDto struct {
	UserId         string  `json:"user_id,omitempty"`
	Address        string  `json:"address,omitempty"`
	Name           string  `json:"name,omitempty"`
	PhoneNumber    string  `json:"phone_number,omitempty"`
	Type           string  `json:"type,omitempty"`
	AddressDefault string  `json:"address_default,omitempty"`
	Lat            float64 `json:"lat,omitempty"`
	Long           float64 `json:"long,omitempty"`
	WardId         string  `json:"ward_id,omitempty"`
	WardText       string  `json:"ward_text,omitempty"`
	DistrictId     string  `json:"district_id,omitempty"`
	DistrictText   string  `json:"district_text,omitempty"`
	ProvinceId     string  `json:"province_id,omitempty"`
	ProvinceText   string  `json:"province_text,omitempty"`
	NationalId     string  `json:"national_id,omitempty"`
	NationalText   string  `json:"national_text,omitempty"`
}

func (c *UsersAddressDto) ToPayload() *users_payload.UsersAddressPayload {
	usersAddressDto := &users_payload.UsersAddressPayload{
		UserId:         c.UserId,
		Address:        c.Address,
		Name:           c.Name,
		PhoneNumber:    c.PhoneNumber,
		Type:           c.Type,
		AddressDefault: c.AddressDefault,
		Lat:            c.Lat,
		Long:           c.Long,
		WardId:         c.WardId,
		WardText:       c.WardText,
		DistrictId:     c.DistrictId,
		DistrictText:   c.DistrictText,
		ProvinceId:     c.ProvinceId,
		ProvinceText:   c.ProvinceText,
		NationalId:     c.NationalId,
		NationalText:   c.NationalText,
	}
	return usersAddressDto
}
