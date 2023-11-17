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
	oauthRouter.GET("/", oauthHandler.RootLoginPage)
	if githubOauthService != nil {
		oauthRouter.GET("/github", oauthHandler.GithubLogin)
		oauthRouter.GET("/github/callback", oauthHandler.GithubLoginCallback)
	}
	if googleOauthService != nil {
		oauthRouter.GET("/google", oauthHandler.GoogleLogin)
		oauthRouter.GET("/google/callback", oauthHandler.GoogleLoginCallback)
	}
}
