package entity

type User struct {
	Name string
	Role string
}

func NewUser() User {
	return User{}
}
