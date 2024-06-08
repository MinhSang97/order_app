package admin_payload

import (
	"encoding/json"
	"github.com/MinhSang97/order_app/model/admin_model"
	"log"
)

type GetAdminRequest struct {
	PassWord    string `json:"-" db:"password, omitempty" validate:"required"`
	Email       string `json:"email,omitempty" db:"email, omitempty" validate:"required"`
	Token       string `json:"token,omitempty" db:"token, omitempty" validate:"required"`
	PhoneNumber string `json:"phone_number,omitempty" db:"phone_number, omitempty" validate:"required"`
}

func (c *GetAdminRequest) ToModel() *admin_model.ReqSignIn {
	admin := &admin_model.ReqSignIn{
		PassWord:    c.PassWord,
		Email:       c.Email,
		Token:       c.Token,
		PhoneNumber: c.PhoneNumber,
	}

	return admin
}

func (c *GetAdminRequest) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
}
