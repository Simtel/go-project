package common

import (
	"github.com/joho/godotenv"
)

func Init() {
	if err := godotenv.Load(".env.local", ".env"); err != nil {
		panic("error load env file")
	}

}
