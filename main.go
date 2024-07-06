package main

import (
	"fmt"
	"time"
)

type User struct {
	name     string
	lastname string
	email    string
}

func (u User) getFullName() string {
	return u.name + " " + u.lastname
}

func main() {
	u := User{name: "Simtel"}

	u.email = "email@example.com"

	fmt.Println(u.name)
	fmt.Println(u.getFullName())
	fmt.Println(time.DateOnly)

	l := Location{name: "Ulyanovsk"}

	fmt.Println(l.getName())
}
