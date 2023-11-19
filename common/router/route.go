package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tae2089/gin-boilerplate/common/domain"
	"github.com/tae2089/gin-boilerplate/common/middleware"
	"github.com/tae2089/gin-boilerplate/common/util"
	"github.com/tae2089/gin-boilerplate/notification"
	"github.com/tae2089/gin-boilerplate/user/repository"
	"gorm.io/gorm"
)

func SetupRouter(e *gin.Engine, db *gorm.DB, jwtKey domain.JwtKey) {

	e.LoadHTMLGlob("templates/*")
	e.Use(middleware.ErrorHandler(notification.GetErrorChan()))
	jwtUtil := util.NewJwtUtil(jwtKey)

	healthRouter := e.Group("/")
	healthRouter.Use(middleware.CheckAccessToken(jwtUtil))
	newHealthRouter(healthRouter)

	userRouter := e.Group("/user")
	userRepository := repository.NewUserRepository(db)
	newUserRouter(userRouter, userRepository, jwtUtil)

	oauthRouter := e.Group("/oauth")
	newOauthRouter(oauthRouter, jwtUtil, userRepository)

}
