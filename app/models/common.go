package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// 自增ID主键
type ID struct {
	ID string `json:"id" gorm:"type:char(64);primaryKey"`
}

func (u *ID) BeforeCreate(db *gorm.DB) (err error) {
	id, err := uuid.NewV4()
	u.ID = id.String()
	return
}

// 创建、更新时间
type Timestamps struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// 软删除
type SoftDeletes struct {
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
