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
// @Accept application/json
// @Produce application/json
// @Param username body string false "用户名"
// @Param password body string false "用户密码"
// @Param email body string false "用户邮箱"
// @Param code body string false "验证码"
// @Param id body string false "验证码对应的 id"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
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

func Info(c *gin.Context) {
	err, user := services.UserService.GetUserInfo(c.Keys["id"].(string))
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, models.UserNotPassword{User: &user})
}

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
