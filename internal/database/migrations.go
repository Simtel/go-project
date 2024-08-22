package database

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"time"
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
				if !tx.Migrator().HasColumn(&contact{}, "Parent") {
					return tx.Migrator().AddColumn(&contact{}, "Parent")
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				type contact struct {
				}
				return tx.Migrator().DropColumn(&contact{}, "Parent")
			},
		},
		{
			ID: "202208301416",
			Migrate: func(tx *gorm.DB) error {
				type domain struct {
					Id        int
					Domain    string
					User      int
					Expired   time.Time
					CreatedAt time.Time
					UpdatedAt time.Time
				}
				return tx.Migrator().CreateTable(&domain{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("domains")
			},
		},
	}
}
