package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/tae2089/gin-boilerplate/common/config"
)

type JwtUtilTestSuit struct {
	suite.Suite
	jwtUtil JwtUtil
}

func TestJwtUtilTestSuite(t *testing.T) {
	suite.Run(t, new(JwtUtilTestSuit))
}

func (s *JwtUtilTestSuit) SetupTest() {
	jwtKey := config.NewJwtKey()
	jwtUtil := NewJwtUtil(jwtKey)
	s.jwtUtil = jwtUtil
}

func (s *JwtUtilTestSuit) TestExmaple1() {
	fmt.Println(1)
}

func (s *JwtUtilTestSuit) TestCreateAccessToken() {
	token, err := s.jwtUtil.CreateAccessToken("abc", false)
	assert.Nil(s.T(), err, nil)
	assert.NotEmpty(s.T(), token.AccessToken)
}

func (s *JwtUtilTestSuit) TestIsAuthorized() {
	token, err := s.jwtUtil.CreateAccessToken("abc", false)
	assert.Nil(s.T(), err, nil)
	result, err := s.jwtUtil.IsAuthorized(token.AccessToken)
	assert.Nil(s.T(), err, nil)
	assert.True(s.T(), result)
}

func (s *JwtUtilTestSuit) TestExtractFieldFromToken() {
	token, err := s.jwtUtil.CreateAccessToken("abc", false)
	assert.Nil(s.T(), err, nil)
	result, err := s.jwtUtil.ExtractFieldFromToken("id", token.AccessToken)
	assert.Nil(s.T(), err, nil)
	assert.Equal(s.T(), "abc", result)
}
