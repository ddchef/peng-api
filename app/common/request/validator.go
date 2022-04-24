package request

import (
	"peng-api/global"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Validator interface {
	GetMessages() ValidatorMessages
}

type ValidatorMessages map[string]string

// GetErrorMsg 获取错误信息
func GetErrorMsg(c *gin.Context, err error) (message string) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		message = err.Error()
	}
	message = removeTopStruct(errs.Translate(global.App.Trans))
	return
}

//   removeTopStruct 定义一个去掉结构体名称前缀的自定义方法：
func removeTopStruct(fileds map[string]string) string {
	rsp := ""
	for _, err := range fileds {
		rsp += err + ";"
	}
	return rsp
}
