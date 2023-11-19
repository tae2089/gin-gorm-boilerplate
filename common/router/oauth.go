package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tae2089/gin-boilerplate/common/util"
	"github.com/tae2089/gin-boilerplate/oauth/handler"
	"github.com/tae2089/gin-boilerplate/oauth/service"
	"github.com/tae2089/gin-boilerplate/user/repository"
)

func newOauthRouter(oauthRouter *gin.RouterGroup, jwtUtil util.JwtUtil, userRepository repository.UserRepository) {
	githubOauth := service.NewGithubOauthProvider()
	googleOauth := service.NewGoogleOauthProvider()
	oauthService := service.NewOauthService(githubOauth, googleOauth, jwtUtil, userRepository)
	oauthHandler := handler.NewOauthHandler(oauthService)

	oauthRouter.GET("/", oauthHandler.RootLoginPage)
	if githubOauth != nil {
		oauthRouter.GET("/github", oauthHandler.GithubLogin)
		oauthRouter.GET("/github/callback", oauthHandler.GithubLoginCallback)
	}
	if googleOauth != nil {
		oauthRouter.GET("/google", oauthHandler.GoogleLogin)
		oauthRouter.GET("/google/callback", oauthHandler.GoogleLoginCallback)
	}
}
