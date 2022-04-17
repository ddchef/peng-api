package services

import (
	"errors"
	"peng-api/app/common/request"
	"peng-api/app/models"
	"peng-api/global"
	"peng-api/utils"
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
	user = models.User{Username: params.Username, Email: params.Email, Password: utils.BcryptMake([]byte(params.Password))}
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
	if user.Active != true {
		err = errors.New("该用户不允许登录")
	}
	return
}

// GetUserInfo 获取用户信息
func (userService *userService) GetUserInfo(id string) (err error, user models.User) {
	err = global.App.DB.First(&user, "id = ?", id).Error
	if err != nil {
		err = errors.New("数据不存在")
	}
	return
}

// GetUserList 获取所有用户列表
func (userService *userService) GetUserList(offset int, limit int) (err error, users []models.UserNotPassword, count int64) {
	err = global.App.DB.Model(&models.User{}).Limit(limit).Offset(offset).Find(&users).Count(&count).Error
	if err != nil {
		err = errors.New("查询失败")
	}
	return
}

func (userService *userService) UpdateUser(id string, params request.BaseUser) (err error, user models.User) {
	err, user = userService.GetUserInfo(id)
	if err != nil {
		return
	}
	err = global.App.DB.Model(&user).Updates(models.User{Username: params.Username, Email: params.Email}).Error
	if err != nil {
		err = errors.New("更新失败")
	}
	return
}

func (userService *userService) DeleteUser(id string) (err error) {
	err, user := userService.GetUserInfo(id)
	if err != nil {
		return
	}
	err = global.App.DB.Delete(&user).Error
	if err != nil {
		err = errors.New("删除失败")
	}
	return
}

func (userService *userService) CreateUser(params request.BaseUser) (err error, user models.User) {
	var result = global.App.DB.Where("username = ?", params.Username).Select("id").First(&models.User{})
	if result.RowsAffected != 0 {
		err = errors.New("用户已存在")
		return
	}
	user = models.User{Username: params.Username, Email: params.Email, Password: utils.BcryptMake([]byte("1qazXSW@@"))}
	err = global.App.DB.Create(&user).Error
	return
}

func (userService *userService) UpdateUserActive(id string, active bool) (err error, user models.User) {
	err, user = userService.GetUserInfo(id)
	if err != nil {
		return
	}
	user.Active = active
	if err = global.App.DB.Save(&user).Error; err != nil {
		err = errors.New("更新状态失败")
	}
	return
}
