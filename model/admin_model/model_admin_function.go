package admin_model

import (
	"encoding/json"
	"log"
	"time"
)

type AdminFunctionModel struct {
	UserId      string    `json:"user_id"  db:"user_id, omitempty"`
	Name        string    `json:"name,omitempty" db:"name, omitempty" validate:"required"`
	Email       string    `json:"email,omitempty" db:"email, omitempty" validate:"required"`
	Password    string    `json:"pass_word,omitempty" db:"password, omitempty" validate:"required"`
	Role        string    `json:"role" db:"role, omitempty"`
	PhoneNumber string    `json:"phone_number,omitempty" db:"phone_number,omitempty" validate:"required"`
	Address     string    `json:"address,omitempty" db:"address,omitempty" validate:"required"`
	CreatedAt   time.Time `json:"-"`
}

func (c *AdminFunctionModel) TableName() string {
	return "admins"
}

func (c *AdminFunctionModel) ToJson() string {
	bs, err := json.Marshal(c)
	if err != nil {
		log.Fatalln(err)

	}
	return string(bs)
}
func (c *AdminFunctionModel) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
}
