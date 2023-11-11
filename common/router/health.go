package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tae2089/gin-boilerplate/common/handler"
)

func newHealthRouter(g *gin.RouterGroup) {
	handler := handler.NewHealthHandler()
	g.GET("/healthz", handler.Healthz)
}
