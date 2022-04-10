package middleware

import (
	"github.com/gin-gonic/gin"
	"peng-api/app/common/response"
	"peng-api/global"
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
		success, err := global.App.Casbin.Enforce(user, "useradmin", path, method)
		if err != nil || success == false {
			response.PermissionsFail(c)
			c.Abort()
			return
		}
	}
}
