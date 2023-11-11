package config

import (
	"os"
	"strconv"

	"github.com/tae2089/gin-boilerplate/common/domain"
	"gopkg.in/gomail.v2"
)

func NewEmailConfig() domain.EmailConfig {
	from := os.Getenv("EMAIL_USER")
	password := os.Getenv("EMAIL_PASSWORD")
	host := os.Getenv("EMAIL_HOST")
	port, err := strconv.Atoi(os.Getenv("EMAIL_PORT"))
	if err != nil {
		panic(err)
	}
	d := gomail.NewDialer(host, port, from, password)
	emailConfig := domain.EmailConfig{From: from, Dialer: d}
	return emailConfig
}
