package payload

type Role int

const (
	ADMIN Role = iota
	EMPLOYEE
	ADMIN1
	ADMIN2
)

func (r Role) String() string {
	return []string{"ADMIN", "EMPLOYEE"}[r]
}
