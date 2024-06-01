package admin_dto

import (
	"github.com/MinhSang97/order_app/payload/admin_payload"
	"time"
)

type AdminFunctionDto struct {
	UserId      string    `json:"user_id"`
	Name        string    `json:"name"  validate:"required"`
	PassWord    string    `json:"pass_word"`
	Email       string    `json:"email" validate:"required,email"`
	Role        string    `json:"role"`
	PhoneNumber string    `json:"-"`
	Address     string    `json:"-"`
	CreatedAt   time.Time `json:"-"`
}

func (c *AdminFunctionDto) ToPayload() *admin_payload.AdminFunctionPayload {
	adminFunctionDto := &admin_payload.AdminFunctionPayload{
		UserId:      c.UserId,
		Name:        c.Name,
		Email:       c.Email,
		Password:    c.PassWord,
		Role:        c.Role,
		PhoneNumber: c.PhoneNumber,
		Address:     c.Address,
		CreatedAt:   c.CreatedAt,
	}

	return adminFunctionDto
}
