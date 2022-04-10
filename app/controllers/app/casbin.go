package app

import (
	"github.com/gin-gonic/gin"
	"peng-api/app/common/request"
	"peng-api/app/common/response"
	"peng-api/app/services"
)

func AddPolicy(c *gin.Context) {
	var policy request.Policy
	if err := c.ShouldBindJSON(&policy); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(policy, err))
		return
	}
	if err := services.CasbinService.AddPolicy(policy); err != nil {
		response.BusinessFail(c, err.Error())
	}
}

func AddRoleForUserInDomain(c *gin.Context) {
	var user request.UserRolePolicy
	if err := c.ShouldBindJSON(&user); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(user, err))
		return
	}
	if err := services.CasbinService.AddRoleForUserInDomain(user); err != nil {
		response.BusinessFail(c, err.Error())
	}
}
