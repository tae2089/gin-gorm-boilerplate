package domain

import "github.com/golang-jwt/jwt/v5"

type JwtKey struct {
	PrivateKey any
	PublicKey  any
	Method     jwt.SigningMethod
}

type JwtCustomClaims struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	jwt.RegisteredClaims
}

type JwtToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
}
