package admin_dto

import (
	"github.com/MinhSang97/order_app/payload/admin_payload"
	"time"
)

type Admin struct {
	UserId      string    `json:"-"`
	Name        string    `json:"name"  validate:"required"`
	PassWord    string    `json:"-" validate:"required"`
	Email       string    `json:"email" validate:"required,email"`
	Role        string    `json:"-"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
	Token       string    `json:"-"`
	PhoneNumber string    `json:"phone_number" validate:"required,number"`
	Address     string    `json:"address" validate:"required"`
}

func (c *Admin) ToPayload() *admin_payload.AddAdminRequest {
	admintPayload := &admin_payload.AddAdminRequest{
		UserId:      c.UserId,
		Name:        c.Name,
		PassWord:    c.PassWord,
		Email:       c.Email,
		Role:        c.Role,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
		Token:       c.Token,
		PhoneNumber: c.PhoneNumber,
		Address:     c.Address,
	}

	return admintPayload
}
