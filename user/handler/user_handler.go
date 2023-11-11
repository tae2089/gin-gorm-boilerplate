package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tae2089/gin-boilerplate/common/util"
	"github.com/tae2089/gin-boilerplate/user/dto"
	"github.com/tae2089/gin-boilerplate/user/service"
)

type UserHandler struct {
	userService service.UserService
	jwtUtil     util.JwtUtil
}

func NewUserHandler(service service.UserService, jwtUtil util.JwtUtil) *UserHandler {
	return &UserHandler{
		userService: service,
		jwtUtil:     jwtUtil,
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
	id, err := u.userService.Login(requestLogin)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	token, err := u.jwtUtil.CreateAccessToken(id, true)
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
