package util

import (
	"github.com/tae2089/bob-logging/logger"
	"github.com/tae2089/gin-boilerplate/common/config"
	"github.com/tae2089/gin-boilerplate/common/domain"
	"gopkg.in/gomail.v2"
)

func SendTextPlainMail(to string, subject string, body string) error {
	err := sendMail([]string{to}, subject, domain.TextPlain, body)
	return err
}

func SendHtmlMail(To string, subject string, body string, attaches ...string) error {
	err := sendMail([]string{To}, subject, domain.HTML, body, attaches...)
	return err
}

func SendBulkTextMail(To []string, subject string, body string, attaches ...string) error {
	err := sendMail(To, subject, domain.TextPlain, body, attaches...)
	return err
}

func SendBulkHTMLMail(To []string, subject string, body string, attaches ...string) error {
	err := sendMail(To, subject, domain.HTML, body, attaches...)
	return err
}

func sendMail(To []string, subject string, bodyType domain.BodyType, body string, attaches ...string) error {
	emailConfig := config.GetEmailConfig()
	m := generateMessage(emailConfig.From, To, subject, bodyType, body, attaches...)
	if err := emailConfig.Dialer.DialAndSend(m); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func generateMessage(from string, To []string, subject string, bodyType domain.BodyType, body string, attaches ...string) *gomail.Message {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", To...)
	m.SetHeader("Subject", subject)
	m.SetBody(string(bodyType), body)
	for _, attach := range attaches {
		m.Attach(attach)
	}
	return m
}
