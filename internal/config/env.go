package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func InitEnv() {
	currentWorkDirectory, _ := os.Getwd()
	fmt.Println("Current Working Directory: " + currentWorkDirectory)

	envLocalPath := filepath.Join(currentWorkDirectory, ".env.local")
	envPath := filepath.Join(currentWorkDirectory, ".env")

	if _, err := os.Stat(envPath); os.IsNotExist(err) {
		rootPath := filepath.Join(currentWorkDirectory, "..", "..", "..", "..")
		envLocalPath = filepath.Join(rootPath, ".env.local")
		envPath = filepath.Join(rootPath, ".env")
	}

	if err := godotenv.Load(envLocalPath, envPath); err != nil {
		panic("error load env file")
	}

}
