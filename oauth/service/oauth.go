package service

import "golang.org/x/oauth2"

type OauthService interface {
	GetAccessToken(code string) (*oauth2.Token, error)
	GetUserInfo(token *oauth2.Token) (string, error)
}

const (
	OAUTH_GOOGLE_URL = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
	OAUTH_GITHUB_URL = "https://api.github.com/user"
)
