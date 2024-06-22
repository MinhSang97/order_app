package req

type ReqSignIn struct {
	Email       string `json:"email,omitempty"`
	PassWord    string `json:"password,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}
