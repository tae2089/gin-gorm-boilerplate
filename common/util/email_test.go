package util

import (
	"testing"

	"github.com/tae2089/gin-boilerplate/common/config"
)

func TestNewEmailUtil(t *testing.T) {
	emailConfig := config.NewEmailConfig()
	emailUtil := NewEmailUtil(emailConfig)
	if emailUtil == nil {
		t.Error("NewEmailUtil error")
	}
}

func TestSendTextPlainMail(t *testing.T) {
	t.Skip("this test is skipped because it is long running")
	emailConfig := config.NewEmailConfig()
	emailUtil := NewEmailUtil(emailConfig)
	err := emailUtil.SendTextPlainMail("test@example.com", "mail testing", "mail testing")
	if err != nil {
		t.Error(err)
	}
}

func TestSendHtmlMail(t *testing.T) {
	t.Skip("this test is skipped because it is long running")
	emailConfig := config.NewEmailConfig()
	emailUtil := NewEmailUtil(emailConfig)
	err := emailUtil.SendHtmlMail("test@example.com", "mail testing", "<h1>mail testing</h1>")
	if err != nil {
		t.Error(err)
	}
}
