package config

import (
	"fmt"
	"os"

	"github.com/tae2089/bob-logging/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var client *gorm.DB

func LoadDBConfig() ConfigOption {
	return func() {
		logger.Info("initializing database")
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Seoul",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
		)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		client = db
		if err != nil {
			panic(err)
		}
		logger.Info("initialized database")
	}
}

func GetDB() *gorm.DB {
	return client
}
