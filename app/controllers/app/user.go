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
		response.ValidateFail(c, request.GetErrorMsg(c, err))
		return
	}

	if err, user := services.UserService.Register(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, models.UserNotPassword{User: &user})
	}
}

// Info 当前用户信息
// @Summary 当前用户信息接口
// @Description 获取当前用户信息接口
// @Tags 用户管理
// @Security BearerAuth
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

// Users 用户列表
// @Summary 用户列表接口
// @Description 获取用户列表接口
// @Tags 用户管理
// @Security BearerAuth
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

// CreateUser 创建用户接口
// @Summary 创建用户接口
// @Tags  用户管理
// @Security BearerAuth
// @Param form body request.BaseUser true "用户信息"
// @Success 200 {object} response.Response{data=models.UserNotPassword}
// @Router /user [post]
func CreateUser(c *gin.Context) {
	var form request.BaseUser
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(c, err))
		return
	}
	if err, user := services.UserService.CreateUser(form); err != nil {
		response.BusinessFail(c, err.Error())
		return
	} else {
		response.Success(c, models.UserNotPassword{User: &user})
	}
}

// DeleteUser 删除用户接口
// @Summary 删除用户接口
// @Tags  用户管理
// @Security BearerAuth
// @Param id path string true "用户ID"
// @Success 200 {object} response.Response
// @Router /user/{id} [delete]
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := services.UserService.DeleteUser(id); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, "")
	}
}

// UpdateUser 更新用户接口
// @Summary 更新用户接口
// @Tags  用户管理
// @Security BearerAuth
// @Param id path string true "用户ID"
// @Param form body request.BaseUser true "用户信息"
// @Success 200 {object} response.Response{data=models.UserNotPassword}
// @Router /user/{id} [put]
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var form request.BaseUser
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(c, err))
		return
	}
	if err, user := services.UserService.UpdateUser(id, form); err != nil {
		response.BusinessFail(c, err.Error())
		return
	} else {
		response.Success(c, models.UserNotPassword{User: &user})
	}
}

// InfoUser 查看用户信息接口
// @Summary 查看用户信息接口
// @Tags  用户管理
// @Security BearerAuth
// @Param id path string true "用户ID"
// @Success 200 {object} response.Response{data=models.UserNotPassword}
// @Router /user/{id} [get]
func InfoUser(c *gin.Context) {
	id := c.Param("id")
	if err, user := services.UserService.GetUserInfo(id); err != nil {
		response.BusinessFail(c, err.Error())
		return
	} else {
		response.Success(c, models.UserNotPassword{User: &user})
	}
}

// UpdateUserActive  更新用户状态接口
// @Summary 更新用户状态接口
// @Tags  用户管理
// @Security BearerAuth
// @Param id path string true "用户ID"
// @Param form body struct{Active bool} true "用户状态"
// @Success 200 {object} response.Response{data=models.UserNotPassword}
// @Router /user/active/:id [put]
func UpdateUserActive(c *gin.Context) {
	id := c.Param("id")
	var form struct {
		Active bool
	}
	if err := c.BindJSON(&form); err != nil {
		response.BusinessFail(c, "缺少状态字段")
	}
	if err, user := services.UserService.UpdateUserActive(id, form.Active); err != nil {
		response.BusinessFail(c, err.Error())
		return
	} else {
		response.Success(c, models.UserNotPassword{User: &user})
	}
}
