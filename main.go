package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/postgres"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"vint.id/goaccounting/database"
	"vint.id/goaccounting/middleware"
	"vint.id/goaccounting/routes"
)

func landing(c *gin.Context) {
	c.HTML(http.StatusOK, "landing/index.tmpl", gin.H{})
}

func main() {
	godotenv.Load(".env")

	if os.Getenv("ENV") != "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}
	gin.SetMode(gin.DebugMode)

	app := gin.Default()
	app.RedirectTrailingSlash = true
	app.HandleMethodNotAllowed = true
	app.NoRoute(routes.NotFound)

	app.Static("/static", "./static")
	app.LoadHTMLGlob("templates/**/*")

	store, err := postgres.NewStore(database.DB, []byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		// handle err
	}
	app.Use(sessions.Sessions("default", store))
	app.Use(middleware.RecoveryMiddleware())
	app.Use(middleware.AuthRedirect())

	app.GET("/", landing)
	app.GET("/login", routes.Login)
	app.POST("/login", routes.Login)
	app.GET("/register", routes.Register)
	app.POST("/register", routes.Register)
	app.POST("/logout", routes.Logout)

	dashboard := app.Group("/dashboard")
	dashboard.Use(middleware.AuthRedirect())
	{
		dashboard.GET("/", routes.Dashboard)
		dashboard.GET("/wallet_management", routes.WalletManagement)
		dashboard.GET("/report", routes.Report)
		dashboard.GET("/settings", routes.Settings)
	}

	transaction := app.Group("/transaction")
	transaction.Use(middleware.AuthRedirect())
	{
		transaction.GET("/", routes.Transaction)
		transaction.GET("/create-transaction-modal", routes.CreateTransactionModal)
		transaction.GET("/new", routes.NewTransaction)
	}

	if err := app.Run(); err != nil {
		panic(err)
	}
}
