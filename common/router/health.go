package router

import (
	"time"

	cache "github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-gonic/gin"
	"github.com/tae2089/gin-boilerplate/common/handler"
)

func newHealthRouter(g *gin.RouterGroup, cacheStore persist.CacheStore) {
	handler := handler.NewHealthHandler()
	g.GET("/healthz", handler.Healthz)
	g.GET("/hello",
		cache.CacheByRequestURI(cacheStore, 2*time.Second),
		func(c *gin.Context) {
			c.String(200, "hello world")
		},
	)
}
