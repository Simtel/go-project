package common

import (
	"os"
)

var VarDir string = "var"

func InitFileStorage() {
	if _, err := os.Stat(VarDir); os.IsNotExist(err) {
		err := os.Mkdir(VarDir, 0777)
		if err != nil {
			panic("Cannot create dir " + VarDir)
		}

	}
}