package util

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/tae2089/bob-logging/logger"
	"github.com/tae2089/gin-boilerplate/common/config"
	"github.com/tae2089/gin-boilerplate/common/domain"
)

func CreateAccessToken(id string, usingRefreshToken bool) (jwtToken domain.JwtToken, err error) {
	jwtKey := config.GetJwtKey()
	refreshToken := ""
	// A usual scenario is to set the expiration time relative to the current time
	claims := getToken(false, id)
	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)
	accessToken, err := token.SignedString(jwtKey.PrivateKey)
	if err != nil {
		return domain.JwtToken{}, err
	}

	if usingRefreshToken {
		refreshClaims := getToken(true, id)
		token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, refreshClaims)
		refreshToken, err = token.SignedString(jwtKey.PrivateKey)
		if err != nil {
			return domain.JwtToken{}, err
		}
	}

	return domain.JwtToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func getToken(isRefresh bool, userID string) domain.JwtCustomClaims {
	expiresAt := time.Now().Add(time.Hour * 24 * 7)
	if isRefresh {
		expiresAt = time.Now().Add(time.Hour * 24 * 365 * 10)
	}
	claims := domain.JwtCustomClaims{
		Name: "jwt-token",
		ID:   userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "tae2089",
			Subject:   "login",
			Audience:  []string{"tae2089"},
		},
	}
	return claims
}

func IsAuthorized(requestToken string) (bool, error) {
	_, err := parseToken(requestToken)
	if err != nil {
		return false, err
	}

	return true, nil
}

func ExtractFieldFromToken(field, requestToken string) (string, error) {
	token, err := parseToken(requestToken)
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return "", fmt.Errorf("Invalid Token")
	}

	return claims[field].(string), nil
}

func parseToken(requestToken string) (*jwt.Token, error) {
	jwtKey := config.GetJwtKey()
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodEd25519); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey.PublicKey, nil
	})

	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return token, nil
}
