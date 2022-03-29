package request

type Register struct {
	Username string `form:"username" json:"username" binding:"required"`
	Mobile   string `form:"mobile" json:"mobile" binding:"mobile"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (register Register) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"username.required": "用户名称不能为空",
		"mobile.required":   "手机号码不能为空",
		"mobile.mobile":     "手机号码格式不正确",
		"password.required": "用户密码不能为空",
	}
}

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (login Login) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"mobile.mobile":     "手机号码格式不正确",
		"password.required": "用户密码不能为空",
	}
}
