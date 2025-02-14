package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load(".env")
}

func AuthRedirect() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		userID := session.Get("UserID")

		if userID != nil {
			if ctx.Request.URL.Path == "/login" || ctx.Request.URL.Path == "/register" || ctx.Request.URL.Path == "/" {
				ctx.Redirect(http.StatusPermanentRedirect, "/dashboard")
				ctx.Abort()
				return
			}
		}

		if userID == nil {
			if ctx.Request.URL.Path != "/login" && ctx.Request.URL.Path != "/register" && ctx.Request.URL.Path != "/" {
				ctx.Redirect(http.StatusPermanentRedirect, "/")
				ctx.Abort()
				return
			}
		}

		ctx.Next()
	}
}
