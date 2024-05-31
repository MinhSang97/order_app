package req

type ReqAdminFunction struct {
	//UserID      string `json:"user_id" validate:"required"`
	Email       string `json:"email,omitempty" validate:"required"`
	Name        string `json:"name,omitempty" validate:"required"`
	Role        string `json:"role,omitempty" validate:"required"`
	PhoneNumber string `json:"phone_number,omitempty" validate:"required"`
	Address     string `json:"address,omitempty" validate:"required"`
}
