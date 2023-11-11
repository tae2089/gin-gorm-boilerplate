package config

import (
	"fmt"
	"os"

	"github.com/tae2089/bob-logging/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDBConfig() *gorm.DB {
	logger.Info("initializing database")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Seoul",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	logger.Info("initialized database")
	return db
}
