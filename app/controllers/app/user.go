package app

import (
	"peng-api/app/common/request"
	"peng-api/app/common/response"
	"peng-api/app/services"

	"github.com/gin-gonic/gin"
)

// Register 用户注册
func Register(c *gin.Context) {
	var form request.Register
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.Register(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, user)
	}
}

func Info(c *gin.Context) {
	err, user := services.UserService.GetUserInfo(c.Keys["id"].(string))
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, user)
}

func Users(c *gin.Context) {
	err, users := services.UserService.GetUserList()
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, users)
}
