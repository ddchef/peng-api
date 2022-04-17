package models

type User struct {
	ID
	Username   string   `json:"username" gorm:"not null;index;comment:用户名称"`
	Email      string   `json:"email" gorm:"null;comment:用户邮箱"`
	Password   string   `json:"password" gorm:"not null;default:'';comment:用户密码"`
	Active     bool     `json:"active" gorm:"not null;default:false;comment:用户状态"`
	Avatar     string   `json:"avatar" gorm:"not null;default:'';comment:用户头像"`
	RealName   string   `json:"realName" gorm:"not null;comment:用户姓名"`
	RoleId     string   `json:"roleId" gorm:"comment:角色 ID"`
	FrontMenus string   `json:"frontMenus" gorm:"comment:个性化菜单"`
	MenusStyle string   `json:"menusStyle" gorm:"comment:菜单风格"`
	Role       []Role   `gorm:"many2many:user_role"`
	Domain     []Domain `gorm:"many2many:user_domain"`
	Timestamps
	SoftDeletes
}

func (user User) GetUid() string {
	return user.ID.ID
}

func (user User) GetUser() string {
	return user.Username
}

type UserNotPassword struct {
	*User
	Password string `json:"-" gorm:"not null;default:'';comment:用户密码"`
}
