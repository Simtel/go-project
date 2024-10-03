package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-project/internal/database"
)

type Migration struct {
	ID string
}

func RunMigrateStatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate-status",
		Short: "Show migration status",
	}

	cmd.RunE = func(cmd *cobra.Command, args []string) error {

		db := database.NewDbMysql()

		var migrations []Migration
		result := db.Find(&migrations)
		if result.Error != nil {
			fmt.Println("Ошибка при получении миграций:", result.Error)
			return nil
		}

		fmt.Println("Исполненные миграции:")
		for _, m := range migrations {
			fmt.Println(m.ID)
		}

		return nil
	}

	return cmd
}
