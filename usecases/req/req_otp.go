package req

type ReqOTP struct {
	Email string `json:"email,omitempty" validate:"required"`
}

type ReqChangePassword struct {
	Email       string `json:"email,omitempty" validate:"required"`
	PassWordNew string `json:"password_new" db:"pass_word_new"`
}
