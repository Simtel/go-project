package models

type User struct {
	Id       uint   `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name     string `gorm:"size:255;not null" json:"name"`
	Lastname string `gorm:"size:255;not null" json:"lastname"`
	Email    string `gorm:"size:255;not null;unique" json:"email"`
}

func (u *User) GetFullName() string {
	return u.Name + " " + u.Lastname
}

func (u *User) GetName() string {
	return u.Name
}
