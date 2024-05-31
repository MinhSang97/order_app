package payload

import (
	"encoding/json"
	"github.com/MinhSang97/order_app/model"
	"log"
	"time"
)

type OtpPayload struct {
	ID          int64     `json:"id" db:"id"`
	UserId      string    `json:"user_id"  db:"user_id, omitempty"`
	PassWordNew string    `json:"password_new" db:"pass_word_new, omitempty" validate:"required"`
	Email       string    `json:"email,omitempty" db:"email, omitempty" validate:"required"`
	Otp         string    `json:"-" db:"role, omitempty"`
	CreatedAt   time.Time `json:"-"`
}

func (c *OtpPayload) ToModel() *model.OtpModel {
	otp := &model.OtpModel{
		ID:          c.ID,
		UserId:      c.UserId,
		PassWordNew: c.PassWordNew,
		Email:       c.Email,
		Otp:         c.Otp,
		CreatedAt:   c.CreatedAt,
	}

	return otp
}

func (c *OtpPayload) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
}
