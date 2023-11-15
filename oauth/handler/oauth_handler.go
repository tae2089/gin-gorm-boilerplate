package handler

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tae2089/gin-boilerplate/oauth/domain"
	oauth "github.com/tae2089/gin-boilerplate/oauth/service"
)

type OauthHandler struct {
	githubService oauth.OauthService
	googleService oauth.OauthService
}

func NewOauthHandler(githubService, googleService oauth.OauthService) *OauthHandler {
	return &OauthHandler{
		githubService: githubService,
		googleService: googleService,
	}
}

func (o *OauthHandler) GithubLoginCallback(c *gin.Context) {
	code := c.Query("code")
	oauthToken, err := o.githubService.GetAccessToken(code)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	userInfo, err := o.githubService.GetUserInfo(oauthToken)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var githubUserInfo domain.GithubUserInfo
	json.Unmarshal([]byte(userInfo), &githubUserInfo)
	// c.SetCookie("access_token", accessToken, 3600, "/", "localhost", false, true)
	c.JSON(201, gin.H{"isSuccess": true, "userInfo": githubUserInfo})
}

func (o *OauthHandler) GoogleLoginCallback(c *gin.Context) {
	code := c.Query("code")
	oauthToken, err := o.googleService.GetAccessToken(code)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	userInfo, err := o.googleService.GetUserInfo(oauthToken)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var googleUserInfo domain.GoogleUserInfo
	json.Unmarshal([]byte(userInfo), &googleUserInfo)
	// c.SetCookie("access_token", accessToken, 3600, "/", "localhost", false, true)
	c.JSON(201, gin.H{"isSuccess": true, "userInfo": googleUserInfo})
}
