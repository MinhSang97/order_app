package users_model

import (
	"encoding/json"
	"log"
)

type ReqUsersSignIn struct {
	Email       string `json:"email,omitempty" validate:"required"`
	PassWord    string `json:"password,omitempty" validate:"required"`
	Token       string `json:"token,omitempty" validate:"required"`
	PhoneNumber string `json:"phone_number,omitempty" validate:"required"`
}

func (c *ReqUsersSignIn) TableName() string {
	return "reqsignin"
}

func (c *ReqUsersSignIn) ToJson() string {
	bs, err := json.Marshal(c)
	if err != nil {
		log.Fatalln(err)

	}
	return string(bs)
}
func (c *ReqUsersSignIn) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
}
