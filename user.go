package main

import "go-project/users"

type User struct {
	name     string
	lastname string
	email    string
	contact  users.Contact
}

func (u *User) getFullName() string {
	return u.name + " " + u.lastname
}
