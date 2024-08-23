package db

import "time"

type Domain struct {
	Id        int `gorm:"primaryKey;autoIncrement"`
	Domain    string
	User      int
	Expired   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
