package app

import (
	"peng-api/app/common/request"
	"peng-api/app/common/response"
	"peng-api/app/services"

	"github.com/gin-gonic/gin"
)

// 添加策略
// Register 添加策略
// @Summary 添加策略接口
// @Tags 权限管理
// @Security BearerAuth
// @Param policy body request.Policy true "策略"
// @Success 200 {object} response.Response
// @Router /casbin/policy [post]
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

// 添加用户、角色和组的策略
// Register 添加用户、角色和组的策略
// @Summary 添加用户、角色和组的策略接口
// @Tags 权限管理
// @Security BearerAuth
// @Param policy body request.UserRolePolicy true "策略"
// @Success 200 {object} response.Response
// @Router /casbin/user [post]
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
