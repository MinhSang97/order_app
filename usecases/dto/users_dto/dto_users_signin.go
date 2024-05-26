package users_dto

import (
	"github.com/MinhSang97/order_app/payload/users_payload"
)

type ReqSignIn struct {
	PassWord string `json:"-" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

//	func (c *ReqSignIn) ToPayload() *admin_payload.GetAdminRequest {
//		reqSignInPayload := &admin_payload.GetAdminRequest{
//			PassWord: c.PassWord,
//			Email:    c.Email,
//		}
//
//		return reqSignInPayload
//	}
func (c *ReqSignIn) ToPayload() *users_payload.GetUsersRequest {
	reqSignInPayload := &users_payload.GetUsersRequest{
		PassWord: c.PassWord,
		Email:    c.Email,
	}

	return reqSignInPayload
}
