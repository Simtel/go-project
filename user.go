package main

type User struct {
	name     string
	lastname string
	email    string
}

func (u *User) getFullName() string {
	return u.name + " " + u.lastname
}
