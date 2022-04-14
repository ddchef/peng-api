package app

import (
	"peng-api/app/common/request"
	"peng-api/app/common/response"
	"peng-api/app/services"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// 登录
// Register 登录
// @Summary 登录接口
// @Description 登录接口
// @Tags 公共接口
// @Param form body request.Login true "login"
// @Success 200 {object} response.Response{data=services.TokenOutPut}
// @Router /public/login [post]
func Login(c *gin.Context) {
	var form request.Login
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.Login(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		tokenData, err, _ := services.JwtService.CreateToken(services.AppGuardName, user)
		if err != nil {
			response.BusinessFail(c, err.Error())
			return
		}
		response.Success(c, tokenData)
	}
}

// 登出
// Register 登出
// @Summary 登出接口
// @Description 登出接口
// @Tags 公共接口
// @Security ApiKeyAuth
// @Success 200 {object} response.Response
// @Router /common/logout [post]
func Logout(c *gin.Context) {
	err := services.JwtService.JoinBlackList(c.Keys["token"].(*jwt.Token))
	if err != nil {
		response.BusinessFail(c, "登出失败")
		return
	}
	response.Success(c, nil)
}
