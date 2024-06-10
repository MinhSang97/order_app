package users_payload

import (
	"encoding/json"
	users "github.com/MinhSang97/order_app/model/users_model"
	"log"
)

type UsersAddressPayload struct {
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

func (c *UsersAddressPayload) ToModel() *users.UsersAddressModel {
	usersAddress := &users.UsersAddressModel{
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

	return usersAddress
}

func (c *UsersAddressPayload) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
}
