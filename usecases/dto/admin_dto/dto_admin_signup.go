package admin_dto

import (
	"github.com/MinhSang97/order_app/payload/admin_payload"
	"time"
)

type Admin struct {
	UserId         string    `json:"-"`
	Name           string    `json:"name" validate:"required"`
	PassWord       string    `json:"-"`
	Email          string    `json:"email" validate:"required"`
	Role           string    `json:"-"`
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
	Token          string    `json:"-"`
	PhoneNumber    string    `json:"phone_number" validate:"required"`
	Address        string    `json:"address"`
	AddressDefault string    `json:"address_default"`
	Telegram       string    `json:"telegram,omitempty" validate:"required"`
	Sex            string    `json:"sex,omitempty"`
	BirthDate      string    `json:"birth_date,omitempty"`
	Lat            float64   `json:"lat,omitempty"`
	Long           float64   `json:"long,omitempty"`
	WardId         string    `json:"ward_id,omitempty"`
	WardText       string    `json:"ward_text,omitempty"`
	DistrictId     string    `json:"district_id,omitempty"`
	DistrictText   string    `json:"district_text,omitempty"`
	ProvinceId     string    `json:"province_id,omitempty"`
	ProvinceText   string    `json:"province_text,omitempty"`
	NationalId     string    `json:"national_id,omitempty"`
	NationalText   string    `json:"national_text,omitempty"`
}

func (c *Admin) ToPayload() *admin_payload.AddAdminRequest {
	admintPayload := &admin_payload.AddAdminRequest{
		UserId:       c.UserId,
		Name:         c.Name,
		PassWord:     c.PassWord,
		Email:        c.Email,
		Role:         c.Role,
		CreatedAt:    c.CreatedAt,
		UpdatedAt:    c.UpdatedAt,
		Token:        c.Token,
		PhoneNumber:  c.PhoneNumber,
		Address:      c.Address,
		Telegram:     c.Telegram,
		Sex:          c.Sex,
		BirthDate:    c.BirthDate,
		Lat:          c.Lat,
		Long:         c.Long,
		WardId:       c.WardId,
		WardText:     c.WardText,
		DistrictId:   c.DistrictId,
		DistrictText: c.DistrictText,
		ProvinceId:   c.ProvinceId,
		ProvinceText: c.ProvinceText,
		NationalId:   c.NationalId,
		NationalText: c.NationalText,
	}

	return admintPayload
}
