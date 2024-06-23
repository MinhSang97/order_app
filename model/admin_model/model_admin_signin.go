package admin_model

import (
	"encoding/json"
	"log"
)

type ReqSignIn struct {
	Email       string `json:"email,omitempty"`
	PassWord    string `json:"password,omitempty" validate:"required"`
	Token       string `json:"token,omitempty" validate:"required"`
	PhoneNumber string `json:"phone_number,omitempty" validate:"required"`
	UserID      string `json:"user_id,omitempty"`
}

func (c *ReqSignIn) TableName() string {
	return "reqsignis"
}

func (c *ReqSignIn) ToJson() string {
	bs, err := json.Marshal(c)
	if err != nil {
		log.Fatalln(err)

	}
	return string(bs)
}
func (c *ReqSignIn) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
}
