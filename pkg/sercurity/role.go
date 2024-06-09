package sercurity

type Role int

const (
	ADMIN Role = iota
	USERS
	DRIVER
	SHOP
)

func (r Role) String() string {
	return []string{"ADMIN", "USERS", "DRIVER", "SHOP"}[r]
}
