// Package models
// descr 模型通用属于与方法
// author fm
// date 2022/11/15 10:46
package models

import (
	"time"

	"github.com/spf13/cast"
)

// BaseModel 基础模型
type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
}

// Exists 通过 ID 判断是否存在
func (class *BaseModel) Exists() bool {
	return class.ID > 0
}

// NotExists 通过 ID 判断是否不存在
func (class *BaseModel) NotExists() bool {
	return class.ID == 0
}

// GetIdString 获取字符串ID
func (class *BaseModel) GetIdString() string {
	return cast.ToString(class.ID)
}

// Timestamps 时间
type Timestamps struct {
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"column:created_at;index;"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"column:updated_at;index;"`
}

type RowsAffected int64

// ToBool rowsAffected 受影响行是否大于 0
// 是否更新、删除成功
func (class *RowsAffected) ToBool() bool {
	return *class > 0
}
