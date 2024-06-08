package req

type ReqUpdateUser struct {
	//Email       string `json:"email,omitempty"`
	Name string `json:"name,omitempty"`
	Sex  string `json:"sex,omitempty"`
	//PhoneNumber string `json:"phone_number,omitempty"`
	BirthDate string `json:"birth_date,omitempty"`
	Telegram  string `json:"telegram,omitempty"`
}
