package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tae2089/gin-boilerplate/oauth/handler"
	"github.com/tae2089/gin-boilerplate/oauth/service"
)

func newOauthRouter(oauthRouter *gin.RouterGroup) {
	githubOauthService := service.NewGithubOauthService()
	googleOauthService := service.NewGoogleOauthService()
	oauthHandler := handler.NewOauthHandler(githubOauthService, googleOauthService)
	if githubOauthService != nil {
		oauthRouter.GET("/github/login/callback", oauthHandler.GithubLoginCallback)
	}
	if googleOauthService != nil {
		oauthRouter.GET("/google/login/callback", oauthHandler.GoogleLoginCallback)
	}
}
