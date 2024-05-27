package users_payload

import (
	"encoding/json"
	users_model "github.com/MinhSang97/order_app/model/users_model"
	"log"

	"time"
)

type AddUsersRequest struct {
	ID          int64     `json:"id" db:"id" `
	UserId      string    `json:"userid"  db:"user_id, omitempty"`
	Name        string    `json:"name,omitempty" db:"name, omitempty"`
	PassWord    string    `json:"-" db:"password, omitempty"`
	Email       string    `json:"email,omitempty" db:"email, omitempty"`
	Role        string    `json:"role" db:"role, omitempty"`
	CreatedAt   time.Time `json:"created_at" db:"created_at, omitempty"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at, omitempty"`
	Token       string    `json:"-" db:"token"`
	PhoneNumber string    `json:"phone_number,omitempty" db:"phone_number, omitempty"`
	Address     string    `json:"address" db:"address, omitempty"`
}

func (c *AddUsersRequest) ToModel() *users_model.Users {
	admin := &users_model.Users{
		ID:          c.ID,
		UserId:      c.UserId,
		Name:        c.Name,
		PassWord:    c.PassWord,
		Email:       c.Email,
		Role:        c.Role,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
		Token:       c.Token,
		PhoneNumber: c.PhoneNumber,
		Address:     c.Address,
	}

	return admin
}

func (c *AddUsersRequest) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
}
