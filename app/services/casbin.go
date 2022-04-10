package services

import (
	"errors"
	"peng-api/app/common/request"
	"peng-api/global"
)

type casbinService struct {
}

var CasbinService = new(casbinService)

func (casbinService *casbinService) AddPolicy(params request.Policy) (err error) {
	res, err := global.App.Casbin.AddPolicy(params.Role, params.Dom, params.Path, params.Method)
	if err != nil {
		err = errors.New("创建策略错误")
	}
	if res != true {
		err = errors.New("创建策略失败，策略已存在")
	}
	return
}

func (casbinService *casbinService) AddRoleForUserInDomain(params request.UserRolePolicy) (err error) {
	res, err := global.App.Casbin.AddRoleForUserInDomain(params.User, params.Role, params.Dom)
	if err != nil {
		err = errors.New("关联角色错误")
	}
	if res != true {
		err = errors.New("关联角色错误失败，已存在")
	}
	return
}
