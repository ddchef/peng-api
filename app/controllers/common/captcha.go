package common

import (
	"peng-api/app/common/response"
	"peng-api/app/services"

	"github.com/gin-gonic/gin"
)

// 获取验证码ID
// Register 获取验证码ID接口
// @Summary 获取验证码ID接口
// @Tags 公共接口
// @Success 200 {object} response.Response{data=string}
// @Router /public/captcha [post]
func CreateCaptcha(c *gin.Context) {
	id, err := services.CaptchaService.CreateCaptcha()
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, id)
	return
}

// 验证码图片
// Register 验证码图片接口
// @Summary 验证码图片接口
// @Tags 公共接口
// @Param id path string true "验证码id"
// @Success 200 {string} string
// @Router /public/captcha/{id} [get]
func CaptchaImage(c *gin.Context) {
	id := c.Param("id")
	captchaImage := services.CaptchaService.CaptchaImage(id)
	response.Image(c, captchaImage)
}
