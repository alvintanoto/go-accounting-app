package routes

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func NotFound(c *gin.Context) {
	isLoggedIn := checkLoggedIn(c)

	if isLoggedIn {
		session := sessions.Default(c)
		userID := session.Get("UserID")
		email := session.Get("Email")

		c.HTML(http.StatusNotFound, "error/index.tmpl", gin.H{
			"code":        http.StatusNotFound,
			"message":     "Page Not Found",
			"description": "The requested page was not found.",
			"isLoggedIn":  isLoggedIn,
			"user_id":     userID,
			"email":       email,
			"sideMenus":   getSideNavigationItem(-1),
		})
		return
	}

	c.HTML(http.StatusNotFound, "error/index.tmpl", gin.H{
		"code":        http.StatusNotFound,
		"message":     "Page Not Found",
		"description": "The requested page was not found.",
	})
}
