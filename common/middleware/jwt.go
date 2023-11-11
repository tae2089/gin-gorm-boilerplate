package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tae2089/bob-logging/logger"
	"github.com/tae2089/gin-boilerplate/common/util"
)

func CheckAccessToken(jwtUtil util.JwtUtil) gin.HandlerFunc {
	return func(c *gin.Context) {
		access_cookie, err := c.Cookie("access_token")
		if errors.Is(err, http.ErrNoCookie) {
			logger.Error(err)
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized API Call",
			})
			c.Abort()
			return
		}
		userID, err := jwtUtil.ExtractFieldFromToken("id", access_cookie)
		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized API Call",
			})
			c.Abort()
			return
		}
		c.Set("user_id", userID)
		c.Next()
	}
}
