package admin_payload

import (
	"encoding/json"
	"github.com/MinhSang97/order_app/model/admin_model"
	"log"
	"time"
)

type AdminFunctionPayload struct {
	UserId      string    `json:"userid"  db:"user_id, omitempty"`
	Name        string    `json:"name,omitempty" db:"name, omitempty"`
	Email       string    `json:"email,omitempty" db:"email, omitempty"`
	Password    string    `json:"pass_word,omitempty" db:"password, omitempty"`
	Role        string    `json:"role" db:"role, omitempty"`
	PhoneNumber string    `json:"phone_number,omitempty" db:"phone_number, omitempty"`
	Address     string    `json:"address" db:"address, omitempty"`
	CreatedAt   time.Time `json:"-"`
}

func (c *AdminFunctionPayload) ToModel() *admin_model.AdminFunctionModel {
	adminFunctionPayload := &admin_model.AdminFunctionModel{
		UserId:      c.UserId,
		Name:        c.Name,
		Email:       c.Email,
		Password:    c.Password,
		Role:        c.Role,
		PhoneNumber: c.PhoneNumber,
		Address:     c.Address,
		CreatedAt:   c.CreatedAt,
	}

	return adminFunctionPayload
}

func (c *AdminFunctionPayload) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
}
