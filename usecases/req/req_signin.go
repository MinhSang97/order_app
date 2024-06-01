package req

type ReqSignIn struct {
	Email    string `json:"email,omitempty" validate:"required"`
	PassWord string `json:"password,omitempty" validate:"required"`
}
