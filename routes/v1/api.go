package v1

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
		publicRouter.POST("/captcha", common.CreateCaptcha)
		publicRouter.GET("/captcha/:id", common.CaptchaImage)

	}
	authRouter.Use(middleware.JWTAuth(services.AppGuardName), middleware.Casbin())
	// 用户管理
	userRouter := authRouter.Group("user")
	{
		userRouter.GET("/info", app.Info)
		userRouter.GET("/users", app.Users)
		userRouter.POST("", app.CreateUser)
		userRouter.PUT("/:id", app.UpdateUser)
		userRouter.DELETE("/:id", app.DeleteUser)
		userRouter.GET("/:id", app.InfoUser)
	}
	// 角色管理
	roleRouter := authRouter.Group("role")
	{
		roleRouter.POST("", app.RoleCreate)
		roleRouter.PUT("/:id", app.RoleUpdate)
		roleRouter.DELETE("/:id", app.RoleDelete)
		roleRouter.GET("/:id", app.RoleInfo)
		roleRouter.GET("/list", app.RoleList)
	}
	// 组管理
	domainRouter := authRouter.Group("domain")
	{
		domainRouter.POST("", app.DomainCreate)
		domainRouter.PUT("/:id", app.DomainUpdate)
		domainRouter.DELETE("/:id", app.DomainDelete)
		domainRouter.GET("/:id", app.DomainInfo)
		domainRouter.GET("/list", app.DomainList)
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
