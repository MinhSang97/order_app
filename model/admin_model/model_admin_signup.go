package admin_model

import (
	"encoding/json"
	"log"
	"time"
)

type Student struct {
	ID           int64     `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Age          int       `json:"age"`
	Grade        float32   `json:"grade"`
	ClassName    string    `json:"class_name"`
	EntranceDate time.Time `json:"-"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
}
type Admin struct {
	ID        int64     `json:"id" db:"id"`
	UserId    string    `json:"-"  db:"user_id, omitempty"`
	Name      string    `json:"name,omitempty" db:"name, omitempty" validate:"required"`
	PassWord  string    `json:"-" db:"password, omitempty" validate:"required"`
	Email     string    `json:"email,omitempty" db:"email, omitempty" validate:"required"`
	Role      string    `json:"-" db:"role, omitempty"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Token     string    `json:"-" db:"token"`
}

func (c *Admin) TableName() string {
	return "admins"
}

func (c *Admin) ToJson() string {
	bs, err := json.Marshal(c)
	if err != nil {
		log.Fatalln(err)

	}
	return string(bs)
}
func (c *Admin) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
}

func (c *Student) TableName() string {
	return "students"
}

func (c *Student) ToJson() string {
	bs, err := json.Marshal(c)
	if err != nil {
		log.Fatalln(err)

	}
	return string(bs)
}

func (c *Student) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
}
