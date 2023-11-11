package util_test

import (
	"testing"

	"github.com/tae2089/gin-boilerplate/common/config"
	"github.com/tae2089/gin-boilerplate/common/util"
)

func TestCreateAccessToken(t *testing.T) {
	jwtKey := config.NewJwtKey()
	jwtUtil := util.NewJwtUtil(jwtKey)
	token, err := jwtUtil.CreateAccessToken("abc", false)
	if err != nil {
		t.Errorf("CreateAccessToken() error = %v", err)
		return
	}
	t.Log(token)

}

func TestIsAuthorized(t *testing.T) {
	jwtKey := config.NewJwtKey()
	jwtUtil := util.NewJwtUtil(jwtKey)
	token, err := jwtUtil.CreateAccessToken("abc", false)
	if err != nil {
		t.Errorf("CreateAccessToken() error = %v", err)
		return
	}
	result, err := jwtUtil.IsAuthorized(token.AccessToken)
	if err != nil {
		t.Errorf("IsAuthorized() error = %v", err)
		return
	}
	t.Log(result)
}

func TestExtractFieldFromToken(t *testing.T) {
	jwtKey := config.NewJwtKey()
	jwtUtil := util.NewJwtUtil(jwtKey)
	token, err := jwtUtil.CreateAccessToken("abc", false)
	if err != nil {
		t.Errorf("CreateAccessToken() error = %v", err)
		return
	}
	result, err := jwtUtil.ExtractFieldFromToken("id", token.AccessToken)
	if err != nil {
		t.Errorf("ExtractFieldFromToken() error = %v", err)
		return
	}
	t.Log(result)
}
