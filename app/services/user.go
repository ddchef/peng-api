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
func (userService *userService) GetUserList() (err error, users []models.UserNotPassword) {
	err = global.App.DB.Find(&users).Error
	if err != nil {
		err = errors.New("查询失败")
	}
	return
}
