package models

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;comment:主键" json:"id,omitempty"`
}

type CommonTimestampsField struct {
	CreatedAt time.Time      `gorm:"not null;comment:创建时间;"`
	UpdatedAt time.Time      `gorm:"not null;comment:更新时间;"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:删除时间;"`
}
