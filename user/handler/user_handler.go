package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tae2089/gin-boilerplate/user/dto"
	"github.com/tae2089/gin-boilerplate/user/service"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (h *UserHandler) Join(c *gin.Context) {
	var requestJoin dto.RequestJoin
	if err := c.ShouldBindJSON(&requestJoin); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if _, err := h.userService.Join(&requestJoin); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"isSuccess": true})
}

func (u *UserHandler) Login(c *gin.Context) {
	requestLogin := &dto.RequestLogin{}
	if err := c.ShouldBindJSON(requestLogin); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	responseLogin, err := u.userService.Login(requestLogin)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie("access_token", responseLogin.AccessToken, 3600, "/", "localhost", false, true)
	if responseLogin.RefreshToken != "" {
		c.SetCookie("refresh_token", responseLogin.RefreshToken, 3600, "/", "localhost", false, true)
	}
	c.JSON(201, gin.H{"isSuccess": true})
}
