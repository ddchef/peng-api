package routes

import (
	"peng-api/app/controllers/app"
	"peng-api/app/controllers/common"
	"peng-api/app/middleware"
	"peng-api/app/services"

	"github.com/gin-gonic/gin"
)

func SetApiGroupRoutes(router *gin.RouterGroup) {
	authRouter := router.Group("")
	publicRouter := router.Group("public")
	{
		publicRouter.POST("/register", app.Register)
		publicRouter.POST("/login", app.Login)
		publicRouter.GET("/captcha", common.CreateCaptcha)
		publicRouter.GET("/captcha/:id", common.CaptchaImage)

	}
	authRouter.Use(middleware.JWTAuth(services.AppGuardName), middleware.Casbin())
	// 用户管理
	userRouter := authRouter.Group("user")
	{
		userRouter.POST("/info", app.Info)
		userRouter.GET("/users", app.Users)
	}
	// 权限管理
	casbinRouter := authRouter.Group("casbin")
	{
		casbinRouter.POST("/policy", app.AddPolicy)
		casbinRouter.POST("/user", app.AddRoleForUserInDomain)
	}
	// 系统
	commonRouter := authRouter.Group("common")
	{
		commonRouter.POST("/logout", app.Logout)
	}
}
