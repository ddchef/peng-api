package middleware

import (
	"peng-api/app/common/response"
	"peng-api/global"

	"github.com/gin-gonic/gin"
)

func Casbin() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _ := c.Get("casbin_user")
		path := c.FullPath()
		method := c.Request.Method
		if user == nil {
			response.PermissionsFail(c)
			c.Abort()
			return
		}
		success, err := global.App.Casbin.Enforce(user, "admin", path, method)
		if err != nil || success == false {
			response.PermissionsFail(c)
			c.Abort()
			return
		}
	}
}
