package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	switch ctx.Request.Method {
	case http.MethodGet:
		ctx.HTML(http.StatusNotFound, "auth/login.tmpl", gin.H{
			"route": ctx.Request.RequestURI,
		})
	}
}

func Register(ctx *gin.Context) {
	switch ctx.Request.Method {
	case http.MethodGet:
		ctx.HTML(http.StatusNotFound, "auth/register.tmpl", gin.H{
			"route": ctx.Request.RequestURI,
		})
	}
}
