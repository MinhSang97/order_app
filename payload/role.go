package payload

type Role int

const (
	ADMIN Role = iota
	USERS
	ADMIN1
	ADMIN2
)

func (r Role) String() string {
	return []string{"ADMIN", "USERS"}[r]
}
