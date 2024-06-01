package req

type ReqAdminFunction struct {
	Email       string `json:"email,omitempty" validate:"required"`
	Name        string `json:"name,omitempty" validate:"required"`
	Role        string `json:"role,omitempty" validate:"required"`
	PhoneNumber string `json:"phone_number,omitempty" validate:"required"`
	Address     string `json:"address,omitempty" validate:"required"`
}

type ReqAdminFunctionAdd struct {
	Email       string `json:"email,omitempty" validate:"required"`
	PassWord    string `json:"pass_word,omitempty" validate:"required"`
	Name        string `json:"name,omitempty" validate:"required"`
	PhoneNumber string `json:"phone_number,omitempty" validate:"required"`
	Address     string `json:"address,omitempty" validate:"required"`
	Role        string `json:"role,omitempty" validate:"required"`
}
