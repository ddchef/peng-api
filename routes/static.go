package routes

import (
	"github.com/gin-gonic/gin"
)

func SetStaticRoutes(router *gin.Engine) {
	// 前端项目静态资源
	router.StaticFile("/", "./view/dist/index.html")
	router.Static("/assets", "./view/dist/assets")
	router.StaticFile("/favicon.ico", "./view/dist/favicon.ico")
	// 其他静态资源
	router.Static("/static", "./static")
	router.Static("/storage", "./static/storage")
}
