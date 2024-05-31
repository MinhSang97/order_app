package model

import (
	"encoding/json"
	"log"
	"time"
)

type OtpModel struct {
	ID          int64     `json:"id" db:"id"`
	UserId      string    `json:"user_id"  db:"user_id, omitempty"`
	PassWordNew string    `json:"password_new" db:"pass_word_new, omitempty" validate:"required"`
	Email       string    `json:"email,omitempty" db:"email, omitempty" validate:"required"`
	Otp         string    `json:"-" db:"role, omitempty"`
	CreatedAt   time.Time `json:"-"`
}

func (c *OtpModel) TableName() string {
	return "admins"
}

func (c *OtpModel) ToJson() string {
	bs, err := json.Marshal(c)
	if err != nil {
		log.Fatalln(err)

	}
	return string(bs)
}
func (c *OtpModel) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
}
