package main

import (
	"fmt"
	"os"
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
	u := User{name: "Simtel", lastname: "Simuls"}

	t := os.Args

	name, lastname := u.name, u.lastname

	u.email = "email@example.com"

	fmt.Println(u.name)
	fmt.Println(u.getFullName())
	fmt.Println(name + " " + lastname)
	fmt.Println(time.DateOnly)

	fmt.Println("args:")
	fmt.Println(len(t))
	location := Location{id: 1, name: "Ulyanovsk"}

	fmt.Println(location.getName())
}
