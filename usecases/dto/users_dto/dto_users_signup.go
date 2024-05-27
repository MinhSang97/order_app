package users_dto

import (
	users_payload "github.com/MinhSang97/order_app/payload/users_payload"
	"time"
)

type Users struct {
	ID          int64     `json:"-"`
	UserId      string    `json:"-"`
	Name        string    `json:"name"  validate:"required"`
	PassWord    string    `json:"-" validate:"required"`
	Email       string    `json:"email" validate:"required,email"`
	Role        string    `json:"-"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
	Token       string    `json:"-"`
	PhoneNumber string    `json:"-"`
	Address     string    `json:"-"`
}

func (c *Users) ToPayload() *users_payload.AddUsersRequest {
	usersPayload := &users_payload.AddUsersRequest{
		ID:          c.ID,
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
	return usersPayload
}
