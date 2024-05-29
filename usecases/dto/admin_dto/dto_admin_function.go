package admin_dto

import (
	"github.com/MinhSang97/order_app/payload/admin_payload"
)

type AdminFunctionDto struct {
	UserId      string `json:"-"`
	Name        string `json:"name"  validate:"required"`
	PassWord    string `json:"-" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Role        string `json:"-"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}

func (c *AdminFunctionDto) ToPayload() *admin_payload.AdminFunctionPayload {
	adminFunctionDto := &admin_payload.AdminFunctionPayload{
		UserId:      c.UserId,
		Name:        c.Name,
		Email:       c.Email,
		Role:        c.Role,
		PhoneNumber: c.PhoneNumber,
		Address:     c.Address,
	}

	return adminFunctionDto
}
