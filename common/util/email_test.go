package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/tae2089/gin-boilerplate/common/config"
)

type EmailUtilTestSuit struct {
	suite.Suite
	emailUtil EmailUtil
}

func TestEmailUtilTestSuite(t *testing.T) {
	suite.Run(t, new(EmailUtilTestSuit))
}

func (s *EmailUtilTestSuit) SetupTest() {
	emailConfig := config.NewEmailConfig()
	emailUtil := NewEmailUtil(emailConfig)
	// emailUtil := NewMockEmailUtil(s.T())

	s.emailUtil = emailUtil
}

func (s *EmailUtilTestSuit) TestSendTextPlainMail() {
	s.T().Skip("this test is skipped because it is long running")
	// s.emailUtil.Mock.On("SendTextPlainMail", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	err := s.emailUtil.SendTextPlainMail("test@example.com", "mail testing", "mail testing")
	assert.Nil(s.T(), err)
}

func (s *EmailUtilTestSuit) TestSendTextPlainMailError() {
	s.T().Skip("this test is skipped because it is long running")
	// s.emailUtil.Mock.On("SendTextPlainMail", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("dial error"))
	err := s.emailUtil.SendTextPlainMail("test@example.com", "mail testing", "mail testing")
	assert.Error(s.T(), err)
}

func (s *EmailUtilTestSuit) TestSendHtmlMail() {
	s.T().Skip("this test is skipped because it is long running")
	// s.emailUtil.Mock.On("SendHtmlMail", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	err := s.emailUtil.SendHtmlMail("test@example.com", "mail testing", "<h1>mail testing</h1>")
	assert.Nil(s.T(), err)
}

func (s *EmailUtilTestSuit) TestSendBulkTextMail() {
	s.T().Skip("this test is skipped because it is long running")
	// s.emailUtil.Mock.On("SendBulkTextMail", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	err := s.emailUtil.SendBulkTextMail([]string{"test@example.com", "test2@example.com"}, "mail testing", "mail testing")
	assert.Nil(s.T(), err)
}

func (s *EmailUtilTestSuit) TestSendBulkHTMLMail() {
	s.T().Skip("this test is skipped because it is long running")
	// s.emailUtil.Mock.On("SendBulkHTMLMail", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	err := s.emailUtil.SendBulkHTMLMail([]string{"test@example.com", "test2@example.com"}, "mail testing", "<h1>mail testing</h1>")
	assert.Nil(s.T(), err)
}
