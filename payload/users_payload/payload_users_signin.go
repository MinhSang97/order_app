package users_payload

import (
	"encoding/json"
	users "github.com/MinhSang97/order_app/model/users_model"
	"log"
)

type GetUsersRequest struct {
	PassWord    string `json:"-" db:"password, omitempty" validate:"required"`
	Email       string `json:"email,omitempty" db:"email, omitempty"`
	Token       string `json:"token,omitempty" validate:"required"`
	PhoneNumber string `json:"phone_number,omitempty" db:"phone_number, omitempty" validate:"required"`
	UserID      string `json:"user_id,omitempty"`
}

func (c *GetUsersRequest) ToModel() *users.ReqUsersSignIn {
	admin := &users.ReqUsersSignIn{
		PassWord:    c.PassWord,
		Email:       c.Email,
		Token:       c.Token,
		PhoneNumber: c.PhoneNumber,
		UserID:      c.UserID,
	}

	return admin
}

func (c *GetUsersRequest) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
}
