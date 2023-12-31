package config

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/joho/godotenv"
)

type ConfigOption func()

func init() {
	projectDir := getProjectDir()
	fmt.Println("loading.env file")
	if os.Getenv("APP_ENV") != "container" {
		err := godotenv.Load(projectDir + "/.env")
		if err != nil {
			log.Panic("Error loading.env file")
		}
	}
}

func getProjectDir() string {
	projectDir := ""
	_, filename, _, _ := runtime.Caller(0)
	projectDir = path.Join(path.Dir(filename), "../../")
	return projectDir
}
