package request

type Policy struct {
	Role   string `form:"role" json:"role" binding:"required"`
	Dom    string `form:"dom" json:"dom" binding:"required"`
	Path   string `form:"path" json:"path" binding:"required"`
	Method string `form:"method" json:"method" binding:"required"`
}

func (policy Policy) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"role.required":   "角色不能为空",
		"dom.required":    "组不能为空",
		"path.required":   "接口不能为空",
		"method.required": "方法不能为空",
	}
}

type UserRolePolicy struct {
	User string `form:"user" json:"user" binding:"required"`
	Role string `form:"role" json:"role" binding:"required"`
	Dom  string `form:"dom" json:"dom" binding:"required"`
}

func (userRolePolicy UserRolePolicy) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"role.required": "角色不能为空",
		"dom.required":  "组不能为空",
		"user.required": "用户不能为空",
	}
}
