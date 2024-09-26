package config

import (
	"github.com/joho/godotenv"
	"os"
	"regexp"
)

const projectDirName = "go-project"

func InitEnv() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))
	if err := godotenv.Load(string(rootPath)+`/.env.local`, string(rootPath)+`/.env`); err != nil {
		panic("error load env file")
	}

}
