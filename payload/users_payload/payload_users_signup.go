package users_payload

import (
	"encoding/json"
	users_model "github.com/MinhSang97/order_app/model/users_model"
	"log"

	"time"
)

type AddUsersRequest struct {
	UserId       string    `json:"userid"  db:"user_id, omitempty"`
	Name         string    `json:"name,omitempty" db:"name, omitempty"`
	PassWord     string    `json:"-" db:"password, omitempty"`
	Email        string    `json:"email,omitempty" db:"email, omitempty"`
	Role         string    `json:"role" db:"role, omitempty"`
	CreatedAt    time.Time `json:"created_at" db:"created_at, omitempty"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at, omitempty"`
	Token        string    `json:"-" db:"token"`
	PhoneNumber  string    `json:"phone_number,omitempty" db:"phone_number, omitempty"`
	Address      string    `json:"address" db:"address, omitempty"`
	Telegram     string    `json:"telegram,omitempty" db:"telegram,omitempty"`
	Sex          string    `json:"sex,omitempty" db:"sex,omitempty"`
	BirthDate    string    `json:"birth_date,omitempty" db:"birth_date,omitempty"`
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

func (c *AddUsersRequest) ToModel() *users_model.Users {
	admin := &users_model.Users{
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

	return admin
}

func (c *AddUsersRequest) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
}
