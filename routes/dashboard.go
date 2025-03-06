package routes

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type sideNavItem struct {
	Name   string
	Icon   string
	Link   string
	Active bool
}

func getSideNavigationItem(active int) []sideNavItem {
	menuPlaceholder := []sideNavItem{
		{
			Name:   "Dashboard",
			Icon:   "bi-house",
			Link:   "/dashboard",
			Active: false,
		},
		{
			Name:   "Transaction",
			Icon:   "bi-cash",
			Link:   "/transaction",
			Active: false,
		},
		{
			Name:   "Wallet Management",
			Icon:   "bi-wallet",
			Link:   "/dashboard/wallet_management",
			Active: false,
		},
		{
			Name:   "Report",
			Icon:   "bi-pie-chart-fill",
			Link:   "/dashboard/report",
			Active: false,
		},
		{
			Name:   "Setting",
			Icon:   "bi-gear",
			Link:   "/dashboard/settings",
			Active: false,
		},
	}

	renderedMenu := []sideNavItem{}
	for i, sideNavItem := range menuPlaceholder {
		sideNavItem.Active = active == i
		renderedMenu = append(renderedMenu, sideNavItem)
	}

	return renderedMenu
}

func Dashboard(ctx *gin.Context) {
	session := sessions.Default(ctx)
	userID := session.Get("UserID")
	email := session.Get("Email")

	ctx.HTML(http.StatusOK, "dashboard/index.tmpl", gin.H{
		"isLoggedIn": checkLoggedIn(ctx),
		"user_id":    userID,
		"email":      email,
		"sideMenus":  getSideNavigationItem(0),
	})
}

func Report(ctx *gin.Context) {
	session := sessions.Default(ctx)
	userID := session.Get("UserID")
	email := session.Get("Email")

	ctx.HTML(http.StatusOK, "report/index.tmpl", gin.H{
		"isLoggedIn": checkLoggedIn(ctx),
		"user_id":    userID,
		"email":      email,
		"sideMenus":  getSideNavigationItem(3),
	})
}

func Settings(ctx *gin.Context) {
	session := sessions.Default(ctx)
	userID := session.Get("UserID")
	email := session.Get("Email")

	ctx.HTML(http.StatusOK, "settings/index.tmpl", gin.H{
		"isLoggedIn": checkLoggedIn(ctx),
		"user_id":    userID,
		"email":      email,
		"sideMenus":  getSideNavigationItem(4),
	})
}
