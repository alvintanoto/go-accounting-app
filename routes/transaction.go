package routes

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Transaction(ctx *gin.Context) {
	session := sessions.Default(ctx)
	userID := session.Get("UserID")
	email := session.Get("Email")

	ctx.HTML(http.StatusOK, "transaction/index.tmpl", gin.H{
		"isLoggedIn": checkLoggedIn(ctx),
		"user_id":    userID,
		"email":      email,
		"sideMenus":  getSideNavigationItem(1),
	})
}

func CreateTransactionModal(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "transaction/create_transaction_modal.tmpl", gin.H{})
}

func NewTransaction(ctx *gin.Context) {
	session := sessions.Default(ctx)
	userID := session.Get("UserID")
	email := session.Get("Email")

	ctx.HTML(http.StatusOK, "transaction/new_transaction.tmpl", gin.H{
		"isLoggedIn": checkLoggedIn(ctx),
		"user_id":    userID,
		"email":      email,
		"sideMenus":  getSideNavigationItem(1),
		"hideClose":  true,
	})
}
