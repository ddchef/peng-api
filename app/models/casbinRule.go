package models

type CasbinRule struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Ptype string `gorm:"size:128;uniqueIndex:unique_index"`
	V0    string `gorm:"size:128;uniqueIndex:unique_index"`
	V1    string `gorm:"size:128;uniqueIndex:unique_index"`
	V2    string `gorm:"size:128;uniqueIndex:unique_index"`
	V3    string `gorm:"size:128;uniqueIndex:unique_index"`
	V4    string `gorm:"size:128;uniqueIndex:unique_index"`
	V5    string `gorm:"size:128;uniqueIndex:unique_index"`
	Name  string `json:"name" gorm:"size:128;comment:策略名称"`
}
