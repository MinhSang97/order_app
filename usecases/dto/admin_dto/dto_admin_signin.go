package admin_dto

import "github.com/MinhSang97/order_app/payload/admin_payload"

type ReqSignIn struct {
	PassWord    string `json:"-"`
	Token       string `json:"-"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

func (c *ReqSignIn) ToPayload() *admin_payload.GetAdminRequest {
	reqSignInPayload := &admin_payload.GetAdminRequest{
		PassWord:    c.PassWord,
		Email:       c.Email,
		Token:       c.Token,
		PhoneNumber: c.PhoneNumber,
	}

	return reqSignInPayload
}
