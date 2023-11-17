package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"io"

	"github.com/tae2089/gin-boilerplate/common/config"
	"golang.org/x/oauth2"
)

type GithubOauth struct {
	oauth2.Config
}

func NewGithubOauthService() OauthService {
	githubOauthConfig := config.GetGithubConfig()
	if githubOauthConfig.ClientID == "" || githubOauthConfig.ClientSecret == "" {
		return nil
	}
	return &GithubOauth{
		githubOauthConfig,
	}
}

func (g *GithubOauth) GetRedirectURL() (redirectURL string, state string) {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	state = base64.URLEncoding.EncodeToString(b)
	redirectURL = g.AuthCodeURL(state)
	return redirectURL, state
}

func (g *GithubOauth) GetAccessToken(code string) (*oauth2.Token, error) {
	token, err := g.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (g *GithubOauth) GetUserInfo(token *oauth2.Token) (string, error) {

	client := g.Client(context.Background(), token)
	resp, err := client.Get(OAUTH_GITHUB_URL)
	if err != nil {
		return "", err
	}
	// Read the response as a byte slice
	respbody, _ := io.ReadAll(resp.Body)

	// Convert byte slice to string and return
	return string(respbody), nil
}
