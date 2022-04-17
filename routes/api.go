package routes

import (
	v1 "peng-api/routes/v1"

	"github.com/gin-gonic/gin"
)

func SetV1ApiGroupRoutes(router *gin.RouterGroup) {
	v1Router := router.Group("/v1")
	v1.SetApiGroupRoutes(v1Router)
}
