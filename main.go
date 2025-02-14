package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/postgres"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"vint.id/goaccounting/database"
	"vint.id/goaccounting/routes"
)

func notFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "error/index.tmpl", gin.H{
		"code":        http.StatusNotFound,
		"message":     "Page Not Found",
		"description": "The requested page was not found.",
	})
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
	app.NoRoute(notFound)

	app.Static("/static", "./static")
	app.LoadHTMLGlob("templates/**/*")

	store, err := postgres.NewStore(database.DB, []byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		// handle err
	}
	app.Use(sessions.Sessions("default", store))

	app.GET("/login", routes.Login)
	app.GET("/register", routes.Register)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
