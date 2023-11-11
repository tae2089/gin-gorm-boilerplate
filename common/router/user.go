package router

import (
	"github.com/gin-gonic/gin"
	userHandler "github.com/tae2089/gin-boilerplate/user/handler"
	"gorm.io/gorm"
)

func newUserRouter(userRouter *gin.RouterGroup, db *gorm.DB) {
	userRouter.POST("/join", userHandler.Join)
	userRouter.POST("/login", userHandler.Login)
}
