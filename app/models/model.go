// Package models
// descr 模型通用属于与方法
// author fm
// date 2022/11/15 10:46
package models

import (
	"time"
)

// BaseModel 基础模型
type BaseModel struct {
	ID uint64 `json:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
}

// Timestamps 时间
type Timestamps struct {
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"column:created_at;index;"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"column:updated_at;index;"`
}
