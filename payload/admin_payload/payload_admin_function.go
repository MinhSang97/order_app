package admin_payload

import (
	"encoding/json"
	"github.com/MinhSang97/order_app/model/admin_model"
	"log"
)

type AdminFunctionPayload struct {
	UserId      string `json:"userid"  db:"user_id, omitempty"`
	Name        string `json:"name,omitempty" db:"name, omitempty"`
	Email       string `json:"email,omitempty" db:"email, omitempty"`
	Role        string `json:"role" db:"role, omitempty"`
	PhoneNumber string `json:"phone_number,omitempty" db:"phone_number, omitempty"`
	Address     string `json:"address" db:"address, omitempty"`
}

func (c *AdminFunctionPayload) ToModel() *admin_model.AdminFunctionModel {
	adminFunctionPayload := &admin_model.AdminFunctionModel{
		UserId:      c.UserId,
		Name:        c.Name,
		Email:       c.Email,
		Role:        c.Role,
		PhoneNumber: c.PhoneNumber,
		Address:     c.Address,
	}

	return adminFunctionPayload
}

func (c *AdminFunctionPayload) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
}
