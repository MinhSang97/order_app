package users_payload

import (
	"encoding/json"
	users "github.com/MinhSang97/order_app/model/users_model"
	"log"
)

type GetUsersRequest struct {
	PassWord string `json:"-" db:"password, omitempty" validate:"required"`
	Email    string `json:"email,omitempty" db:"email, omitempty" validate:"required"`
	Token    string `json:"token,omitempty" validate:"required"`
}

func (c *GetUsersRequest) ToModel() *users.ReqUsersSignIn {
	admin := &users.ReqUsersSignIn{
		PassWord: c.PassWord,
		Email:    c.Email,
		Token:    c.Token,
	}

	return admin
}

func (c *GetUsersRequest) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
}
