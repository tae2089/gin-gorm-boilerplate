package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"

	"github.com/tae2089/bob-logging/logger"
	"github.com/tae2089/gin-boilerplate/common/config"
	"golang.org/x/oauth2"
)

type GoogleOauth struct {
	oauth2.Config
}

func NewGoogleOauthProvider() OauthProvider {
	googleOauthConfig := config.GetGoogleConfig()
	if googleOauthConfig.ClientID == "" || googleOauthConfig.ClientSecret == "" || googleOauthConfig.RedirectURL == "" {
		logger.Error(errors.New("google oauth config error"))
		return nil
	}
	return &GoogleOauth{
		googleOauthConfig,
	}
}

func (g *GoogleOauth) GetRedirectURL() (redirectURL string, state string) {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	state = base64.URLEncoding.EncodeToString(b)
	redirectURL = g.AuthCodeURL(state)
	return redirectURL, state
}

func (g *GoogleOauth) GetAccessToken(code string) (*oauth2.Token, error) {
	token, err := g.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (g *GoogleOauth) GetUserInfo(token *oauth2.Token) ([]byte, error) {

	client := g.Client(context.Background(), token)
	resp, err := client.Get(OAUTH_GOOGLE_URL + token.AccessToken)
	if err != nil {
		return nil, err
	}
	// Read the response as a byte slice
	respbody, _ := io.ReadAll(resp.Body)
	// Convert byte slice to string and return
	return respbody, nil
}
