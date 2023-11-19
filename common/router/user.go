package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tae2089/gin-boilerplate/common/util"
	"github.com/tae2089/gin-boilerplate/user/handler"
	"github.com/tae2089/gin-boilerplate/user/repository"
	"github.com/tae2089/gin-boilerplate/user/service"
)

func newUserRouter(userRouter *gin.RouterGroup, userRepository repository.UserRepository, jwtUtil util.JwtUtil) {
	userService := service.NewUserService(userRepository, jwtUtil)
	userHandler := handler.NewUserHandler(userService)
	userRouter.POST("/join", userHandler.Join)
	userRouter.POST("/login", userHandler.Login)
}
