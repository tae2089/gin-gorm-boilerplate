package util

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/tae2089/bob-logging/logger"
	"github.com/tae2089/gin-boilerplate/common/domain"
)

type JwtUtil interface {
	CreateAccessToken(id string, usingRefreshToken bool) (jwtToken domain.JwtToken, err error)
	ExtractFieldFromToken(field, requestToken string) (string, error)
	IsAuthorized(requestToken string) (bool, error)
}

type jwtUtil struct {
	domain.JwtKey
}

func NewJwtUtil(key domain.JwtKey) JwtUtil {
	return &jwtUtil{key}
}

func (j *jwtUtil) CreateAccessToken(id string, usingRefreshToken bool) (jwtToken domain.JwtToken, err error) {
	refreshToken := ""
	// A usual scenario is to set the expiration time relative to the current time
	claims := j.getToken(false, id)
	claims.Subject = "access_token"
	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)
	accessToken, err := token.SignedString(j.PrivateKey)
	if err != nil {
		return domain.JwtToken{}, err
	}

	if usingRefreshToken {
		refreshClaims := j.getToken(true, id)
		refreshClaims.Subject = "refresh_token"
		token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, refreshClaims)
		refreshToken, err = token.SignedString(j.PrivateKey)
		if err != nil {
			return domain.JwtToken{}, err
		}
	}

	return domain.JwtToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (j *jwtUtil) getToken(isRefresh bool, userID string) domain.JwtCustomClaims {
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
		},
	}
	return claims
}

func (j *jwtUtil) IsAuthorized(requestToken string) (bool, error) {
	_, err := j.parseToken(requestToken)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (j *jwtUtil) ExtractFieldFromToken(field, requestToken string) (string, error) {
	token, err := j.parseToken(requestToken)
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return "", fmt.Errorf("Invalid Token")
	}

	return claims[field].(string), nil
}

func (j *jwtUtil) parseToken(requestToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodEd25519); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return j.PublicKey, nil
	})

	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return token, nil
}
