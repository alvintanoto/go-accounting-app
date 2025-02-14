package routes

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Dashboard(ctx *gin.Context) {
	session := sessions.Default(ctx)
	userID := session.Get("UserID")
	email := session.Get("Email")

	ctx.HTML(http.StatusOK, "dashboard/index.tmpl", gin.H{
		"isLoggedIn": checkLoggedIn(ctx),
		"user_id":    userID,
		"email":      email,
	})
}
