package util

import (
	"github.com/tae2089/bob-logging/logger"
	"github.com/tae2089/gin-boilerplate/common/domain"
	"gopkg.in/gomail.v2"
)

type EmailUtil interface {
	SendTextPlainMail(to string, subject string, body string) error
	SendHtmlMail(To string, subject string, body string, attaches ...string) error
	SendBulkTextMail(To []string, subject string, body string, attaches ...string) error
	SendBulkHTMLMail(To []string, subject string, body string, attaches ...string) error
}

type emailUtil struct {
	domain.EmailConfig
}

func NewEmailUtil(emailConfig domain.EmailConfig) EmailUtil {
	return &emailUtil{emailConfig}
}

func (e *emailUtil) SendTextPlainMail(to string, subject string, body string) error {
	err := e.sendMail([]string{to}, subject, domain.TextPlain, body)
	return err
}

func (e *emailUtil) SendHtmlMail(To string, subject string, body string, attaches ...string) error {
	err := e.sendMail([]string{To}, subject, domain.HTML, body, attaches...)
	return err
}

func (e *emailUtil) SendBulkTextMail(To []string, subject string, body string, attaches ...string) error {
	err := e.sendMail(To, subject, domain.TextPlain, body, attaches...)
	return err
}

func (e *emailUtil) SendBulkHTMLMail(To []string, subject string, body string, attaches ...string) error {
	err := e.sendMail(To, subject, domain.HTML, body, attaches...)
	return err
}

func (e *emailUtil) sendMail(To []string, subject string, bodyType domain.BodyType, body string, attaches ...string) error {
	m := e.generateMessage(e.From, To, subject, bodyType, body, attaches...)
	if err := e.Dialer.DialAndSend(m); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (e *emailUtil) generateMessage(from string, To []string, subject string, bodyType domain.BodyType, body string, attaches ...string) *gomail.Message {
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
