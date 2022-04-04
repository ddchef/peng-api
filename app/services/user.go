package services

import (
	"errors"
	"peng-api/app/common/request"
	"peng-api/app/models"
	"peng-api/global"
	"peng-api/utils"
	"strconv"
)

type userService struct {
}

var UserService = new(userService)

// Register 注册
func (userService *userService) Register(params request.Register) (err error, user models.User) {
	verify := CaptchaService.Verify(params.ID, params.Code)
	if verify != true {
		err = errors.New("验证码错误")
		return
	}
	var result = global.App.DB.Where("username = ?", params.Username).Select("id").First(&models.User{})
	if result.RowsAffected != 0 {
		err = errors.New("用户名已存在")
		return
	}
	user = models.User{Username: params.Username, Mobile: params.Mobile, Password: utils.BcryptMake([]byte(params.Password))}
	err = global.App.DB.Create(&user).Error
	return
}

// Login 登录
func (userService *userService) Login(params request.Login) (err error, user *models.User) {
	verify := CaptchaService.Verify(params.Id, params.Code)
	if verify != true {
		err = errors.New("验证码错误")
		return
	}
	err = global.App.DB.Where("username = ?", params.Username).First(&user).Error
	if err != nil || !utils.BcryptMakeCheck([]byte(params.Password), user.Password) {
		err = errors.New("用户名不存在或密码错误")
	}
	return
}

// GetUserInfo 获取用户信息
func (userService *userService) GetUserInfo(id string) (err error, user models.User) {
	intId, err := strconv.Atoi(id)
	err = global.App.DB.First(&user, intId).Error
	if err != nil {
		err = errors.New("数据不存在")
	}
	return
}

// 获取所有用户列表
func (userService *userService) GetUserList(offset int, limit int) (err error, users []models.UserNotPassword, count int64) {
	err = global.App.DB.Model(&models.User{}).Limit(limit).Offset(offset).Find(&users).Count(&count).Error
	if err != nil {
		err = errors.New("查询失败")
	}
	return
}
