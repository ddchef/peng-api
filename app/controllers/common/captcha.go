package common

import (
	"peng-api/app/common/response"
	"peng-api/app/services"

	"github.com/gin-gonic/gin"
)

func CreateCaptcha(c *gin.Context) {
	id, err := services.CaptchaService.CreateCaptcha()
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, id)
	return
}

func CaptchaImage(c *gin.Context) {
	id := c.Param("id")
	captchaImage := services.CaptchaService.CaptchaImage(id)
	response.Image(c, captchaImage)
}
