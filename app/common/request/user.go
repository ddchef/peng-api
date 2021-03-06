package request

type BaseUser struct {
	Username string `form:"username" json:"username" binding:"min=4,max=10,required"`
	Email    string `form:"email" json:"email" binding:"email"`
	RealName string `form:"realName" json:"realName" binding:"min=4,max=10,required"`
}

// func (user BaseUser) GetMessages() ValidatorMessages {
// 	return ValidatorMessages{
// 		"username.required": "用户名称不能为空",
// 		"username.min":      "用户名称长度为4~10",
// 		"username.max":      "用户名称长度为4~10",
// 		"email.required":    "邮箱不能为空",
// 		"email.email":       "邮箱格式不正确",
// 		"realName.required": "真实姓名不能为空",
// 	}
// }

type Register struct {
	BaseUser
	Password string `form:"password" json:"password" binding:"required"`
	Code     string `form:"code" json:"code" binding:"required"`
	ID       string `form:"id" json:"id" binding:"required"`
}

// func (register Register) GetMessages() ValidatorMessages {
// 	return ValidatorMessages{
// 		"username.required": "用户名称不能为空",
// 		"email.required":    " 邮箱不能为空",
// 		"email.email":       " 邮箱格式不正确",
// 		"password.required": "用户密码不能为空",
// 		"code.required":     "验证码不能为空",
// 		"id.required":       "验证码不能为空",
// 	}
// }

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Id       string `form:"id" json:"id" binding:"required"`
	Code     string `form:"code" json:"code" binding:"required"`
}

// func (login Login) GetMessages() ValidatorMessages {
// 	return ValidatorMessages{
// 		"username.required": "用户名不能为空",
// 		"password.required": "用户密码不能为空",
// 		"code.required":     "验证码不能为空",
// 		"id.required":       "验证码不能为空",
// 	}
// }
