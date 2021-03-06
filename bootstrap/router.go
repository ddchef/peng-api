package bootstrap

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"peng-api/global"
	"peng-api/routes"
	"syscall"
	"time"

	docs "peng-api/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	routes.SetStaticRoutes(router)

	// 注册 api 分组路由
	apiGroup := router.Group("/api")
	routes.SetV1ApiGroupRoutes(apiGroup)

	return router
}

// RunServer 启动服务器
func RunServer() {
	r := setupRouter()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", global.App.Config.App.Port),
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
