package routes

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func WalletManagement(ctx *gin.Context) {
	session := sessions.Default(ctx)
	userID := session.Get("UserID")
	email := session.Get("Email")

	ctx.HTML(http.StatusOK, "wallet_management/index.tmpl", gin.H{
		"isLoggedIn": checkLoggedIn(ctx),
		"user_id":    userID,
		"email":      email,
		"sideMenus":  getSideNavigationItem(2),
	})
}

// get content.
