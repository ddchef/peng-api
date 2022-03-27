package routes

import (
	"peng-api/app/controllers/app"
	"peng-api/app/middleware"
	"peng-api/app/services"

	"github.com/gin-gonic/gin"
)

func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.POST("/user/register", app.Register)
	router.POST("/auth/login", app.Login)
	authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	authRouter.POST("/auth/info", app.Info)
	authRouter.GET("/auth/users", app.Users)
	authRouter.POST("/auth/logout", app.Logout)
}
