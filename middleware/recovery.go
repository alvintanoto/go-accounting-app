package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Next()
		defer func() {
			if err := recover(); err != nil {
				c.HTML(http.StatusInternalServerError, "error/index.tmpl", gin.H{
					"code":        http.StatusInternalServerError,
					"message":     "Internal Server Error",
					"description": "An unexpected error occured, try again later.",
				})
			}
			c.Abort()
		}()
	}
}
