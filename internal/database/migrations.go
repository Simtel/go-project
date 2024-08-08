package database

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func GetMigrations() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		{
			ID: "202208301400",
			Migrate: func(tx *gorm.DB) error {
				type contact struct {
					ID   uint `gorm:"primaryKey;uniqueIndex"`
					Name string
				}
				return tx.Migrator().CreateTable(&contact{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("contacts")
			},
		},
		{
			ID: "202208301415",
			Migrate: func(tx *gorm.DB) error {
				type contact struct {
					Parent int
				}
				return tx.Migrator().AddColumn(&contact{}, "Parent")
			},
			Rollback: func(tx *gorm.DB) error {
				type contact struct {
					Age int
				}
				return tx.Migrator().DropColumn(&contact{}, "Parent")
			},
		},
	}
}
