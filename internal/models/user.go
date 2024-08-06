package models

type User struct {
	Id       uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Name     string `gorm:"size:255;not null"`
	Lastname string `gorm:"size:255;not null"`
	Email    string `gorm:"size:255;not null;unique"`
}

func (u *User) GetFullName() string {
	return u.Name + " " + u.Lastname
}

func (u *User) GetName() string {
	return u.Name
}
