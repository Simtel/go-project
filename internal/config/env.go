package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
)

const projectDirName = "go-project"

func InitEnv() {
	currentWorkDirectory, _ := os.Getwd()
	fmt.Println("Current Working Directory: " + currentWorkDirectory)

	// Попробуем загрузить .env из текущей директории, затем из корня проекта
	envLocalPath := filepath.Join(currentWorkDirectory, ".env.local")
	envPath := filepath.Join(currentWorkDirectory, ".env")

	// Если в текущей директории нет .env, ищем в корне проекта (4 уровня вверх)
	if _, err := os.Stat(envPath); os.IsNotExist(err) {
		rootPath := filepath.Join(currentWorkDirectory, "..", "..", "..", "..")
		envLocalPath = filepath.Join(rootPath, ".env.local")
		envPath = filepath.Join(rootPath, ".env")
	}

	if err := godotenv.Load(envLocalPath, envPath); err != nil {
		panic("error load env file")
	}

}
