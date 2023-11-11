package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tae2089/gin-boilerplate/common/config"
	"github.com/tae2089/gin-boilerplate/common/router"
	"github.com/tae2089/gin-boilerplate/user/model"
)

func main() {
	e := gin.Default()
	config.LoadingConfigOptions(
		config.LoadDBConfig(),
		config.LoadJwtConfig(),
		// config.LoadEmailConfig(),
	)
	db := config.GetDB()
	db.AutoMigrate(&model.User{})
	router.SetupRouter(e, db)
	e.Run()
}
