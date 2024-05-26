package admin_model

import (
	"encoding/json"
	"log"
	"time"
)

type Admin struct {
	//ID          int64     `json:"id" db:"id"`
	UserId      string    `json:"-"  db:"user_id, omitempty"`
	Name        string    `json:"name,omitempty" db:"name, omitempty" validate:"required"`
	PassWord    string    `json:"-" db:"pass_word, omitempty" validate:"required"`
	Email       string    `json:"email,omitempty" db:"email, omitempty" validate:"required"`
	Role        string    `json:"-" db:"role, omitempty"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
	Token       string    `json:"-" db:"token"`
	PhoneNumber string    `json:"phone_number,omitempty" db:"phone_number,omitempty" validate:"required"`
	Address     string    `json:"address,omitempty" db:"address,omitempty" validate:"required"`
}

func (c *Admin) TableName() string {
	return "admins"
}

func (c *Admin) ToJson() string {
	bs, err := json.Marshal(c)
	if err != nil {
		log.Fatalln(err)

	}
	return string(bs)
}
func (c *Admin) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
}
