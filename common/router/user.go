package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tae2089/gin-boilerplate/common/util"
	"github.com/tae2089/gin-boilerplate/user/handler"
	"github.com/tae2089/gin-boilerplate/user/repository"
	"github.com/tae2089/gin-boilerplate/user/service"
	"gorm.io/gorm"
)

func newUserRouter(userRouter *gin.RouterGroup, db *gorm.DB, jwtUtil util.JwtUtil) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, jwtUtil)
	userHandler := handler.NewUserHandler(userService)
	userRouter.POST("/join", userHandler.Join)
	userRouter.POST("/login", userHandler.Login)
}
