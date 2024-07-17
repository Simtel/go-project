package models

import "go-project/users"

type User struct {
	Name     string
	Lastname string
	Email    string
	Contact  users.Contact
}

func (u *User) GetFullName() string {
	return u.Name + " " + u.Lastname
}

func (u *User) GetName() string {
	return u.Name
}
