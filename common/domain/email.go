package domain

import "gopkg.in/gomail.v2"

type EmailConfig struct {
	From   string
	Dialer *gomail.Dialer
}

type BodyType string

const (
	TextPlain BodyType = "text/plain"
	HTML      BodyType = "text/html"
)
