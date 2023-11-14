package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tae2089/gin-boilerplate/oauth/handler"
	"github.com/tae2089/gin-boilerplate/oauth/service"
)

func newOauthRouter(oauthRouter *gin.RouterGroup) {
	githubOauthService := service.NewGithubOauthService()
	oauthHandler := handler.NewOauthHandler(githubOauthService)
	if githubOauthService != nil {
		oauthRouter.GET("/github/login/callback", oauthHandler.GithubLoginCallback)
	}
}
