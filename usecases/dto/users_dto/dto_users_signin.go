package users_dto

import (
	"github.com/MinhSang97/order_app/payload/users_payload"
)

type ReqSignIn struct {
	PassWord string `json:"-" validate:"required"`
	Token    string `json:"-"`
	Email    string `json:"email" validate:"required,email"`
}

func (c *ReqSignIn) ToPayload() *users_payload.GetUsersRequest {
	reqSignInPayload := &users_payload.GetUsersRequest{
		PassWord: c.PassWord,
		Email:    c.Email,
		Token:    c.Token,
	}

	return reqSignInPayload
}
