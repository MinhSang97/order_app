package users_model

import (
	"encoding/json"
	"log"
	"time"
)

type Users struct {
	ID        int64     `json:"id" db:"id"`
	UserId    string    `json:"-"  db:"user_id, omitempty"`
	Name      string    `json:"name,omitempty" db:"name, omitempty" validate:"required"`
	PassWord  string    `json:"-" db:"password, omitempty" validate:"required"`
	Email     string    `json:"email,omitempty" db:"email, omitempty" validate:"required"`
	Role      string    `json:"-" db:"role, omitempty"`
	CreatedAt time.Time `json:"-" db:"created_at, omitempty"`
	UpdatedAt time.Time `json:"-" db:"updated_at, omitempty"`
	Token     string    `json:"-" db:"token"`
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
