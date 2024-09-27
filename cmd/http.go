package cmd

import (
	"github.com/spf13/cobra"
	"go-project/internal/app"
	"log"
	"net/http"
)

func RunHTTP() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "http",
		Short: "Run http server",
	}

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		a, _ := app.GetApplication()

		// добавление обработчиков маршрутов
		a.Container().AddHandler(a.Container().GetDomainsApi())
		a.Container().AddHandler(a.Container().GetMainApi())
		a.Container().AddHandler(a.Container().GetUsersApi())

		log.Println("Server is starting on port 3000...")
		err := http.ListenAndServe(":3000", a.Container().GetRouter())
		if err != nil {
			log.Fatalf("Server failed: %v", err)
		}
		return nil
	}

	return cmd
}
