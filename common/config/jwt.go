package config

import (
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/tae2089/gin-boilerplate/common/domain"
)

func NewJwtKey() domain.JwtKey {
	var jwtKey domain.JwtKey
	privateKeyPath := os.Getenv("PRIVATE_KEY_PATH")
	publicKeyPath := os.Getenv("PUBLIC_KEY_PATH")

	privateKeyBytes, err := os.ReadFile(privateKeyPath)
	if err != nil {
		log.Fatal(err)
	}
	// 개인 키 파싱
	block, _ := pem.Decode(privateKeyBytes)
	if block == nil || block.Type != "PRIVATE KEY" {
		log.Fatal("Failed to decode PEM block containing RSA private key")
	}

	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	publicKeyBytes, err := os.ReadFile(publicKeyPath)
	if err != nil {
		log.Fatal(err)
	}
	block, _ = pem.Decode(publicKeyBytes)
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	jwtKey.PrivateKey = key
	jwtKey.PublicKey = pub
	jwtKey.Method = getSigingMethod()
	return jwtKey
}

func getSigingMethod() jwt.SigningMethod {
	method := os.Getenv("JWT_SIGNING_METHOD")

	switch method {
	case "HS384":
		return jwt.SigningMethodHS384
	case "HS512":
		return jwt.SigningMethodHS512
	case "ES256":
		return jwt.SigningMethodES256
	case "ES384":
		return jwt.SigningMethodES384
	case "ES512":
		return jwt.SigningMethodES512
	case "RS256":
		return jwt.SigningMethodRS256
	case "RS384":
		return jwt.SigningMethodRS384
	case "RS512":
		return jwt.SigningMethodRS512
	case "PS256":
		return jwt.SigningMethodPS256
	case "PS384":
		return jwt.SigningMethodPS384
	case "PS512":
		return jwt.SigningMethodPS512
	case "EdDSA":
		return jwt.SigningMethodEdDSA
	default:
		return jwt.SigningMethodNone
	}
}
