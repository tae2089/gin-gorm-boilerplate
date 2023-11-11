package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tae2089/gin-boilerplate/common/domain"
	"github.com/tae2089/gin-boilerplate/common/middleware"
	"github.com/tae2089/gin-boilerplate/common/util"
	"gorm.io/gorm"
)

func SetupRouter(e *gin.Engine, db *gorm.DB, jwtKey domain.JwtKey) {

	jwtUtil := util.NewJwtUtil(jwtKey)

	// config.LoadEnv()
	healthRouter := e.Group("/")
	healthRouter.Use(middleware.CheckAccessToken(nil))
	newHealthRouter(healthRouter)

	userRouter := e.Group("/user")
	newUserRouter(userRouter, db, jwtUtil)
}
