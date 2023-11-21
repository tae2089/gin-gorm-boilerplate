package router

import (
	"time"

	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tae2089/gin-boilerplate/common/domain"
	"github.com/tae2089/gin-boilerplate/common/middleware"
	"github.com/tae2089/gin-boilerplate/common/util"
	"github.com/tae2089/gin-boilerplate/notification"
	"github.com/tae2089/gin-boilerplate/user/repository"
	"gorm.io/gorm"
)

func SetupRouter(e *gin.Engine, db *gorm.DB, jwtKey domain.JwtKey, cacheStore persist.CacheStore) {

	e.LoadHTMLGlob("templates/*")
	useCorsMiddleware(e)
	useErrorSenderMiddleware(e)
	jwtUtil := util.NewJwtUtil(jwtKey)

	healthRouter := e.Group("/")
	newHealthRouter(healthRouter, cacheStore)

	userRouter := e.Group("/user")
	userRepository := repository.NewUserRepository(db)
	newUserRouter(userRouter, userRepository, jwtUtil)

	oauthRouter := e.Group("/oauth")
	newOauthRouter(oauthRouter, jwtUtil, userRepository)

}

func useCorsMiddleware(e *gin.Engine) {
	e.Use(cors.New(
		cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"*"},
			AllowHeaders:     []string{"*"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))
}

func useErrorSenderMiddleware(e *gin.Engine) {
	e.Use(middleware.ErrorHandler(notification.GetErrorChan()))
}
