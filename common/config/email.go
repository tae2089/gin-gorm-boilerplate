package config

import (
	"os"
	"strconv"

	"github.com/tae2089/gin-boilerplate/common/domain"
	"gopkg.in/gomail.v2"
)

var emailConfig domain.EmailConfig

func LoadEmailConfig() ConfigOption {
	return func() {
		from := os.Getenv("EMAIL_USER")
		password := os.Getenv("EMAIL_PASSWORD")
		host := os.Getenv("EMAIL_HOST")
		port, err := strconv.Atoi(os.Getenv("EMAIL_PORT"))
		if err != nil {
			panic(err)
		}
		d := gomail.NewDialer(host, port, from, password)
		emailConfig = domain.EmailConfig{From: from, Dialer: d}
	}
}

func GetEmailConfig() domain.EmailConfig {
	return emailConfig
}
