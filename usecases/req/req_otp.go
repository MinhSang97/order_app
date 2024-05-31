package req

type ReqOTP struct {
	Email string `json:"email,omitempty" validate:"required"`
}
