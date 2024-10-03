package main

import (
	"github.com/spf13/cobra"
	"go-project/cmd"
	"go-project/internal/app"
	"log"
)

func initApp() {
	a, err := app.NewApp()
	if err != nil {
		log.Fatal("Fail to create app: ", err)
	}

	app.SetApplication(a)
}

// основной запуск приложения
func main() {
	rootCmd := &cobra.Command{}

	cobra.OnInitialize(initApp)

	rootCmd.AddCommand(
		cmd.RunHTTP(),
		cmd.RunMigrate(),
		cmd.RunMigrateStatus(),
	)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("failed to execute root cmd: %v", err)

		return
	}
}
