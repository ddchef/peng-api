package request

type Role struct {
	RoleName          string `form:"roleName" json:"roleName" binding:"max=10,min=3,required"`
	ParentId          string `form:"parentId" json:"parentId" binding:"required"`
	DefaultFrontMenus string `form:"defaultFrontMenus" json:"defaultFrontMenus" binding:"required"`
}

func (role Role) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"roleName.max":               "角色名称长度必须大于10",
		"roleName.min":               "角色名称长度必须大于3",
		"parentId.required":          "父级角色必填项",
		"defaultFrontMenus.required": "默认菜单树必填项",
	}
}

type UpdateRole struct {
	Role
	ID uint `form:"id" json:"id" binding:"required"`
}

func (role UpdateRole) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"roleName.max":               "角色名称长度必须大于10",
		"roleName.min":               "角色名称长度必须大于3",
		"parentId.required":          "父级角色必填项",
		"defaultFrontMenus.required": "默认菜单树必填项",
		"id.required":                "缺失 id",
	}
}
