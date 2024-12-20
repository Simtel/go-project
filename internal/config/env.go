package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

const projectDirName = "go-project"

func InitEnv() {
	currentWorkDirectory, _ := os.Getwd()
	fmt.Println("Current Working Directory: " + currentWorkDirectory)
	if err := godotenv.Load(string(currentWorkDirectory)+`/.env.local`, string(currentWorkDirectory)+`/.env`); err != nil {
		panic("error load env file")
	}

}
