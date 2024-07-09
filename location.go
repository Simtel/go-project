package main

type Location struct {
	id     int
	name   string
	parent *Location
}

func (l *Location) getName() string {
	return l.name
}

func (l *Location) setParent(parent *Location) {
	l.parent = parent
}
