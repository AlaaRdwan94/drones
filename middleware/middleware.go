package middleware

import (
	"task/middleware/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

//SetMiddlewareAuthentication check for the validity of the authentication token provided
func SetMiddlewareAuthentication(next gin.HandlerFunc) gin.HandlerFunc {
	return func (c *gin.Context) {
		serial ,err := auth.ExtractTokenMetadata(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			return
		}
		c.Writer.Header().Set("serial",serial)
		next(c)
	}
}
