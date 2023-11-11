package config

import (
	"fmt"
	"log"
	"path"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/tae2089/bob-logging/logger"
)

type ConfigOption func()

func init() {
	logger.Info("initializing config")
	projectDir := getProjectDir()
	fmt.Println("loading.env file")
	err := godotenv.Load(projectDir + "/.env")
	if err != nil {
		log.Panic("Error loading.env file")
	}
	logger.Info("initalized config")
}

func getProjectDir() string {
	projectDir := ""
	_, filename, _, _ := runtime.Caller(0)
	projectDir = path.Join(path.Dir(filename), "../../")
	return projectDir
}

func LoadingDefaultConfigOptions() {
	for _, option := range []ConfigOption{LoadDBConfig(), LoadJwtConfig()} {
		option()
	}
}

func LoadingConfigOptions(options ...ConfigOption) {
	for _, option := range options {
		option()
	}
}
