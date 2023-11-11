package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tae2089/gin-boilerplate/common/util"
	"github.com/tae2089/gin-boilerplate/user/dto"
	userService "github.com/tae2089/gin-boilerplate/user/service"
)

func Join(c *gin.Context) {
	var requestJoin dto.RequestJoin
	if err := c.ShouldBindJSON(&requestJoin); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if _, err := userService.Join(&requestJoin); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"isSuccess": true})
}

func Login(c *gin.Context) {
	requestLogin := &dto.RequestLogin{}
	if err := c.ShouldBindJSON(requestLogin); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	id, err := userService.Login(requestLogin)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	token, err := util.CreateAccessToken(id, true)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie("access_token", token.AccessToken, 3600, "/", "localhost", false, true)
	if token.RefreshToken != "" {
		c.SetCookie("refresh_token", token.RefreshToken, 3600, "/", "localhost", false, true)
	}
	c.JSON(201, gin.H{
		"success": true,
	})
}
