package main

type Location struct {
	id   int
	name string
}

func (l *Location) getName() string {
	return "Location" + ": " + l.name
}
