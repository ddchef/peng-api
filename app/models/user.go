package models

import "strconv"

type User struct {
	ID
	Username string `json:"username" gorm:"not null;index;comment:用户名称" example:"用户名称"`
	Email    string `json:"email" gorm:"null;comment:用户邮箱" example:"用户邮箱"`
	Password string `json:"password" gorm:"not null;default:'';comment:用户密码" example:"用户密码"`
	Active   bool   `json:"active" gorm:"not null;default:false;comment:用户状态" example:"false"`
	Avatar   string `json:"avatar" gorm:"not null;default:'';comment:用户头像" example:"用户头像"`
	RealName string `json:"realName" gorm:"not null;comment:用户姓名" example:"用户姓名"`
	Timestamps
	SoftDeletes
}

func (user User) GetUid() string {
	return strconv.Itoa(int(user.ID.ID))
}

func (user User) GetUser() string {
	return user.Username
}

type UserNotPassword struct {
	*User
	Password string `json:"-" gorm:"not null;default:'';comment:用户密码"`
}
