package models

type Location struct {
	Id     int
	Name   string
	Parent *Location
}

func (l *Location) GetName() string {
	return l.Name
}

func (l *Location) SetParent(parent *Location) {
	l.Parent = parent
}
