package users_dto

import (
	users_payload "github.com/MinhSang97/order_app/payload/users_payload"
	"time"
)

type Users struct {
	UserId       string    `json:"-"`
	Name         string    `json:"name" validate:"required"`
	PassWord     string    `json:"-"`
	Email        string    `json:"email" validate:"required"`
	Role         string    `json:"-"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
	Token        string    `json:"-"`
	PhoneNumber  string    `json:"-"`
	Address      string    `json:"-"`
	Telegram     string    `json:"telegram,omitempty"`
	Sex          string    `json:"sex,omitempty"`
	BirthDate    string    `json:"birth_date,omitempty"`
	Lat          float64   `json:"lat,omitempty"`
	Long         float64   `json:"long,omitempty"`
	WardId       string    `json:"ward_id,omitempty"`
	WardText     string    `json:"ward_text,omitempty"`
	DistrictId   string    `json:"district_id,omitempty"`
	DistrictText string    `json:"district_text,omitempty"`
	ProvinceId   string    `json:"province_id,omitempty"`
	ProvinceText string    `json:"province_text,omitempty"`
	NationalId   string    `json:"national_id,omitempty"`
	NationalText string    `json:"national_text,omitempty"`
}

func (c *Users) ToPayload() *users_payload.AddUsersRequest {
	usersPayload := &users_payload.AddUsersRequest{
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
		BirthDate:    c.BirthDate,
		Sex:          c.Sex,
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
	return usersPayload
}
