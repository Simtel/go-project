package cmd

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/spf13/cobra"
	"go-project/internal/database"
	"log"
)

func RunMigrate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Migrate database",
	}

	cmd.RunE = func(cmd *cobra.Command, args []string) error {

		db := database.NewDbMysql()
		database.MigrateDB(db)

		m := gormigrate.New(db, gormigrate.DefaultOptions, database.GetMigrations())

		if err := m.Migrate(); err != nil {
			log.Fatalf("Migration failed: %v", err)
		}

		return nil
	}

	return cmd
}
