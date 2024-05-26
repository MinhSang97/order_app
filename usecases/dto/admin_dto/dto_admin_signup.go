package admin_dto

import (
	"github.com/MinhSang97/order_app/payload/admin_payload"
	"time"
)

type Admin struct {
	ID        int64     `json:"-"`
	UserId    string    `json:"-"`
	Name      string    `json:"name"  validate:"required"`
	PassWord  string    `json:"-" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Role      string    `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Token     string    `json:"-"`
}

func (c *Admin) ToPayload() *admin_payload.AddAdminRequest {
	admintPayload := &admin_payload.AddAdminRequest{
		ID:        c.ID,
		UserId:    c.UserId,
		Name:      c.Name,
		PassWord:  c.PassWord,
		Email:     c.Email,
		Role:      c.Role,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
		Token:     c.Token,
	}

	return admintPayload
}
