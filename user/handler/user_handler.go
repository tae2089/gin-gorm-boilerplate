package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tae2089/bob-logging/logger"
	"github.com/tae2089/gin-boilerplate/exception"
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
		logger.Error(err)
		_ = c.Error(exception.CustomError{
			StatusCode: http.StatusBadRequest,
			Code:       exception.INVALID_ERROR,
			Message:    "data is invaild",
		})
		return
	}
	if _, err := h.userService.Join(&requestJoin); err != nil {
		logger.Error(err)
		_ = c.Error(exception.CustomError{
			StatusCode: http.StatusBadRequest,
			Code:       exception.INVALID_ERROR,
			Message:    "failed to save user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"isSuccess": true})
}

func (u *UserHandler) Login(c *gin.Context) {
	requestLogin := &dto.RequestLogin{}
	if err := c.ShouldBindJSON(requestLogin); err != nil {
		logger.Error(err)
		_ = c.Error(exception.CustomError{
			StatusCode: http.StatusBadRequest,
			Code:       exception.INVALID_ERROR,
			Message:    "data is invaild",
		})
		return
	}
	responseLogin, err := u.userService.Login(requestLogin)
	if err != nil {
		logger.Error(err)
		_ = c.Error(exception.CustomError{
			StatusCode: http.StatusBadRequest,
			Code:       exception.INVALID_ERROR,
			Message:    "failed to login",
		})
		return
	}
	c.SetCookie("access_token", responseLogin.AccessToken, 3600, "/", "localhost", false, true)
	if responseLogin.RefreshToken != "" {
		c.SetCookie("refresh_token", responseLogin.RefreshToken, 3600, "/", "localhost", false, true)
	}
	c.JSON(http.StatusAccepted, gin.H{"isSuccess": true})
}
