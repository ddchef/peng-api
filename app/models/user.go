package models

import "strconv"

type User struct {
	ID
	Username string `json:"username" gorm:"not null;index;comment:用户名称"`
	Mobile   string `json:"mobile" gorm:"null;comment:用户手机号"`
	Password string `json:"password" gorm:"not null;default:'';comment:用户密码"`
	Timestamps
	SoftDeletes
}

func (user User) GetUid() string {
	return strconv.Itoa(int(user.ID.ID))
}

type UserNotPassword struct {
	*User
	Password string `json:"-" gorm:"not null;default:'';comment:用户密码"`
}
