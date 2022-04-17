package services

import (
	"errors"
	"peng-api/app/common/request"
	"peng-api/app/models"
	"peng-api/global"
)

type roleService struct{}

var RoleService = new(roleService)

func (roleService *roleService) Create(parmas request.Role) (err error, role models.Role) {
	result := global.App.DB.Where("roleName = ?", parmas.RoleName).Select("id").First(&models.Role{})
	if result.RowsAffected != 0 {
		err = errors.New("角色已存在")
		return
	}
	role = models.Role{RoleName: parmas.RoleName, ParentId: parmas.ParentId, DefaultFrontMenus: parmas.DefaultFrontMenus}
	err = global.App.DB.Create(&role).Error
	return
}

func (roleService *roleService) Update(id string, parmas request.UpdateRole) (err error, role models.Role) {
	err, role = roleService.Info(id)
	if err != nil {
		return
	}
	role.RoleName = parmas.RoleName
	role.ParentId = parmas.ParentId
	role.DefaultFrontMenus = parmas.DefaultFrontMenus
	err = global.App.DB.Save(&role).Error
	if err != nil {
		err = errors.New("数据更新失败")
	}
	return
}

func (roleService *roleService) Info(id string) (err error, role models.Role) {
	err = global.App.DB.First(&role, "id = ?", id).Error
	if err != nil {
		err = errors.New("数据不存在")
	}
	return
}

func (roleService *roleService) List(offset int, limit int) (err error, roles []models.Role, count int64) {
	err = global.App.DB.Model(&models.Role{}).Offset(offset).Limit(limit).Find(&roles).Count(&count).Error
	if err != nil {
		err = errors.New("查询失败")
	}
	return
}

func (roleService *roleService) Delete(id string) (err error) {
	err, role := roleService.Info(id)
	if err != nil {
		return
	}
	err = global.App.DB.Delete(&role).Error
	if err != nil {
		err = errors.New("删除失败")
	}
	return
}
