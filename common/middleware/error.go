package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tae2089/gin-boilerplate/exception"
)

func ErrorHandler(errorEventChanel *chan error) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, err := range c.Errors {
			switch e := err.Err.(type) {
			case exception.CustomError:
				c.AbortWithStatusJSON(e.StatusCode, e)
			default:
				*errorEventChanel <- err
				c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"message": "Service Unavailable"})
			}
		}
	}
}
