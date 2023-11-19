package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	oauth "github.com/tae2089/gin-boilerplate/oauth/service"
)

type OauthHandler struct {
	oauthService *oauth.OauthService
}

func NewOauthHandler(oauthService *oauth.OauthService) *OauthHandler {
	return &OauthHandler{
		oauthService: oauthService,
	}
}

func (o *OauthHandler) RootLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "oauth.html", gin.H{})
}

func (o *OauthHandler) GithubLogin(c *gin.Context) {
	redirectURL := o.oauthService.GithubLogin()
	c.Redirect(http.StatusTemporaryRedirect, redirectURL)
}

func (o *OauthHandler) GoogleLogin(c *gin.Context) {
	redirectURL, state := o.oauthService.GoogleLogin()
	c.SetCookie("state", state, 3600, "/", "localhost", false, true)
	c.Redirect(301, redirectURL)
}

func (o *OauthHandler) GithubLoginCallback(c *gin.Context) {
	code := c.Query("code")
	token, err := o.oauthService.GithubLoginCallback(code)
	if err != nil {
		c.JSON(500, gin.H{"error": "server internal error"})
		return
	}
	c.SetCookie("access_token", token.AccessToken, 3600, "/", "localhost", false, true)
	if token.RefreshToken != "" {
		c.SetCookie("refresh_token", token.RefreshToken, 3600, "/", "localhost", false, true)
	}
	c.JSON(201, gin.H{"isSuccess": true})
}

func (o *OauthHandler) GoogleLoginCallback(c *gin.Context) {
	// if you don't want to use state, please comment 59-64 line
	stateCookie, err := c.Cookie("state")
	stateForm := c.Request.FormValue("state")
	if stateCookie != stateForm {
		c.JSON(400, gin.H{"error": "state not match"})
		return
	}
	code := c.Query("code")
	token, err := o.oauthService.GoogleLoginCallback(code)
	if err != nil {
		c.JSON(500, gin.H{"error": "server internal error"})
		return
	}
	c.SetCookie("access_token", token.AccessToken, 3600, "/", "localhost", false, true)
	if token.RefreshToken != "" {
		c.SetCookie("refresh_token", token.RefreshToken, 3600, "/", "localhost", false, true)
	}
	c.JSON(201, gin.H{"isSuccess": true})
}
