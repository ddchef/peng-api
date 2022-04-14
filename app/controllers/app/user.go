package app

import (
	"peng-api/app/common/request"
	"peng-api/app/common/response"
	"peng-api/app/do"
	"peng-api/app/models"
	"peng-api/app/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

//用户注册
// Register 用户注册接口
// @Summary 用户注册接口
// @Description 可以注册用户
// @Tags 公共接口
// @Param form body request.Register true "用户信息"
// @Success 200 {object} response.Response{data=models.User}
// @Router /public/register [post]
func Register(c *gin.Context) {
	var form request.Register
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.Register(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, models.UserNotPassword{User: &user})
	}
}

// 获取当前用户信息
// Register 当前用户信息
// @Summary 当前用户信息接口
// @Description 获取当前用户信息接口
// @Tags 用户管理
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{data=models.UserNotPassword}
// @Router /user/info [get]
func Info(c *gin.Context) {
	err, user := services.UserService.GetUserInfo(c.Keys["id"].(string))
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, models.UserNotPassword{User: &user})
}

// 用户列表
// Register 用户列表
// @Summary 用户列表接口
// @Description 获取用户列表接口
// @Tags 用户管理
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{data=do.List{list=[]models.UserNotPassword}}
// @Router /user/users [get]
func Users(c *gin.Context) {
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	err, users, count := services.UserService.GetUserList(offset, limit)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, do.List{
		List:  users,
		Total: count,
	})
}
