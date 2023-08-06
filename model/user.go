package model

type User struct {
	Name     string
	Password string
	Email    string
}

func CreateUser(name, password, email string) *User {
	return &User{name, password, email}
}
