package models

type Role struct {
	ID
	RoleName          string `json:"roleName" gorm:"not null;unique;comment:角色名称"`
	ParentId          string `json:"parentId" gorm:"comment:父级角色 id"`
	DefaultFrontMenus string `json:"defaultFrontMenus" gorm:"comment:默认当前角色前端展示菜单树"`
	Timestamps
	SoftDeletes
}
