package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/joho/godotenv"
	"vint.id/goaccounting/models"
	"vint.id/goaccounting/service"
)

func init() {
	godotenv.Load(".env")
}

func checkLoggedIn(ctx *gin.Context) bool {
	session := sessions.Default(ctx)
	userID := session.Get("UserID")
	if userID == nil {
		return false
	}

	return true
}

func createSession(ctx *gin.Context, user *models.User) {
	session := sessions.Default(ctx)
	session.Set("UserID", user.Id)
	session.Set("Email", user.Email)
	session.Save()
}

func error(c *gin.Context, code int, message, description string) {
	c.HTML(http.StatusNotFound, "error/index.tmpl", gin.H{
		"code":        code,
		"message":     message,
		"description": description,
	})
}

func Login(ctx *gin.Context) {
	switch ctx.Request.Method {
	case http.MethodGet:
		ctx.HTML(http.StatusOK, "auth/login.tmpl", gin.H{
			"route": ctx.Request.RequestURI,
		})
	case http.MethodPost:
		var userLogin models.Login
		if err := ctx.Request.ParseForm(); err != nil {
			error(ctx, http.StatusBadRequest, "Bad Request", "Unable to parse form.")
			return
		}
		if err := ctx.ShouldBindWith(&userLogin, binding.Form); err != nil {
			error(ctx, http.StatusBadRequest, err.Error(), "-")
			return
		}

		user, err := service.Login(&userLogin)
		if err != nil {
			fmt.Println(err)
			error(ctx, http.StatusInternalServerError, "Internal Server Error", "Failed to login.")
			return
		}

		createSession(ctx, user)
		ctx.Redirect(http.StatusMovedPermanently, "/dashboard")
	}
}

func Register(ctx *gin.Context) {
	switch ctx.Request.Method {
	case http.MethodGet:
		ctx.HTML(http.StatusOK, "auth/register.tmpl", gin.H{
			"route": ctx.Request.RequestURI,
		})
	case http.MethodPost:
		var user models.User
		if err := ctx.Request.ParseForm(); err != nil {
			error(ctx, http.StatusBadRequest, "Bad Request", "Unable to parse form.")
			return
		}
		if err := ctx.ShouldBindWith(&user, binding.Form); err != nil {
			error(ctx, http.StatusBadRequest, err.Error(), "-")
			return
		}

		err := service.Register(&user)
		if err != nil {
			error(ctx, http.StatusInternalServerError, "Internal Server Error", "Failed to Register user.")
			return
		}

		createSession(ctx, &user)
		ctx.Redirect(http.StatusMovedPermanently, "/dashboard")
	}
}

func Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	session.Save()
	ctx.Redirect(http.StatusMovedPermanently, "/")
}
