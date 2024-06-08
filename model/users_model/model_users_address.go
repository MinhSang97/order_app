package users_model

import (
	"encoding/json"
	"log"
)

type UsersAddressModel struct {
	UserId         string `json:"user_id,omitempty"`
	Address        string `json:"address,omitempty"`
	Name           string `json:"name,omitempty"`
	PhoneNumber    string `json:"phone_number,omitempty"`
	Type           string `json:"type,omitempty"`
	AddressDefault string `json:"address_default,omitempty"`
	Lat            string `json:"lat,omitempty"`
	Long           string `json:"long,omitempty"`
	WardId         string `json:"ward_id,omitempty"`
	WardText       string `json:"ward_text,omitempty"`
	DistrictId     string `json:"district_id,omitempty"`
	DistrictText   string `json:"district_text,omitempty"`
	ProvinceId     string `json:"province_id,omitempty"`
	ProvinceText   string `json:"province_text,omitempty"`
	NationalId     string `json:"national_id,omitempty"`
	NationalText   string `json:"national_text,omitempty"`
}

func (c *UsersAddressModel) TableName() string {
	return "users_address"
}

func (c *UsersAddressModel) ToJson() string {
	bs, err := json.Marshal(c)
	if err != nil {
		log.Fatalln(err)

	}
	return string(bs)
}
func (c *UsersAddressModel) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
}
