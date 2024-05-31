package dto

import (
	"github.com/MinhSang97/order_app/payload"
	"time"
)

type OtpDto struct {
	ID          int64     `json:"id" db:"id"`
	UserId      string    `json:"user_id"  db:"user_id, omitempty"`
	PassWordNew string    `json:"pass_word_new" db:"pass_word_new"`
	Email       string    `json:"email,omitempty" db:"email, omitempty" validate:"required"`
	Otp         string    `json:"-" db:"role, omitempty"`
	CreatedAt   time.Time `json:"-"`
}

func (c *OtpDto) ToPayload() *payload.OtpPayload {
	otpPayload := &payload.OtpPayload{
		ID:          c.ID,
		UserId:      c.UserId,
		PassWordNew: c.PassWordNew,
		Email:       c.Email,
		Otp:         c.Otp,
		CreatedAt:   c.CreatedAt,
	}
	return otpPayload
}
