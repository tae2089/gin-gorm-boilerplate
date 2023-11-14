package config

import (
	"os"

	"github.com/tae2089/gin-boilerplate/common/domain"
)

func GetGithubConfig() domain.GitHubOauthConfig {
	return domain.GitHubOauthConfig{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
	}
}
