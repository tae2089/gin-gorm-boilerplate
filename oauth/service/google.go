package service

import (
	"context"
	"io"

	"github.com/tae2089/gin-boilerplate/common/config"
	"golang.org/x/oauth2"
)

type GoogleOauth struct {
	oauth2.Config
}

func NewGoogleOauthService() OauthService {
	googleOauthConfig := config.GetGoogleConfig()
	if googleOauthConfig.ClientID == "" || googleOauthConfig.ClientSecret == "" {
		return nil
	}
	return &GithubOauth{
		googleOauthConfig,
	}
}

func (g *GoogleOauth) GetAccessToken(code string) (*oauth2.Token, error) {
	token, err := g.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (g *GoogleOauth) GetUserInfo(token *oauth2.Token) (string, error) {

	client := g.Client(context.Background(), token)
	resp, err := client.Get(OAUTH_GOOGLE_URL + token.AccessToken)
	if err != nil {
		return "", err
	}
	// Read the response as a byte slice
	respbody, _ := io.ReadAll(resp.Body)

	// Convert byte slice to string and return
	return string(respbody), nil
}
