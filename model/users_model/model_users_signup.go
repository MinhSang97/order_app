package users_model

import (
	"encoding/json"
	"log"
	"time"
)

type Users struct {
	UserId       string    `json:"-"  db:"user_id, omitempty"`
	Name         string    `json:"name,omitempty" db:"name, omitempty" validate:"required"`
	PassWord     string    `json:"-" db:"password, omitempty" validate:"required"`
	Email        string    `json:"email,omitempty" db:"email, omitempty" validate:"required"`
	Role         string    `json:"-" db:"role, omitempty"`
	CreatedAt    time.Time `json:"-" db:"created_at, omitempty"`
	UpdatedAt    time.Time `json:"-" db:"updated_at, omitempty"`
	Token        string    `json:"-" db:"token"`
	PhoneNumber  string    `json:"phone_number,omitempty" db:"phone_number,omitempty" validate:"required"`
	Address      string    `json:"address,omitempty" db:"address,omitempty" validate:"required"`
	Telegram     string    `json:"telegram,omitempty" db:"telegram,omitempty"`
	Sex          string    `json:"sex,omitempty" db:"sex,omitempty" validate:"required"`
	BirthDate    string    `json:"birth_date,omitempty" db:"birth_date,omitempty" validate:"required"`
	Lat          float64   `json:"lat,omitempty" db:"lat,omitempty"`
	Long         float64   `json:"long,omitempty" db:"long,omitempty"`
	WardId       string    `json:"ward_id,omitempty" db:"ward_id,omitempty"`
	WardText     string    `json:"ward_text,omitempty" db:"ward_text,omitempty"`
	DistrictId   string    `json:"district_id,omitempty" db:"district_id,omitempty"`
	DistrictText string    `json:"district_text,omitempty" db:"district_text,omitempty"`
	ProvinceId   string    `json:"province_id,omitempty" db:"province_id,omitempty"`
	ProvinceText string    `json:"province_text,omitempty" db:"province_text,omitempty"`
	NationalId   string    `json:"national_id,omitempty" db:"national_id,omitempty"`
	NationalText string    `json:"national_text,omitempty" db:"national_text,omitempty"`
}

func (c *Users) TableName() string {
	return "admins"
}

func (c *Users) ToJson() string {
	bs, err := json.Marshal(c)
	if err != nil {
		log.Fatalln(err)

	}
	return string(bs)
}
func (c *Users) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
}
