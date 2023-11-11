package config_test

import (
	"testing"

	"github.com/tae2089/gin-boilerplate/common/config"
)

func TestLoadJwtKey(t *testing.T) {
	jwtKey := config.NewJwtKey()
	if jwtKey.PrivateKey == nil && jwtKey.PublicKey == nil && jwtKey.Method == nil {
		t.Fail()
	}
}
