package req

type ReqSignUp struct {
	Email    string `json:"email,omitempty" validate:"required"`
	PassWord string `json:"password,omitempty" validate:"required"`
	Name     string `json:"name,omitempty" validate:"required"`
}
