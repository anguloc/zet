// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTask = "task"

// Task mapped from table <task>
type Task struct {
	ID            uint64     `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id"`
	Mark          string     `gorm:"column:mark;type:varchar(30);not null;comment:任务标识" json:"mark"`                                    // 任务标识
	Status        uint32     `gorm:"column:status;type:tinyint unsigned;not null;comment:任务状态 0待执行,1执行中,20执行完成,99执行失败" json:"status"`   // 任务状态 0待执行,1执行中,20执行完成,99执行失败
	Data          *string    `gorm:"column:data;type:text;comment:任务数据" json:"data"`                                                    // 任务数据
	Result        *string    `gorm:"column:result;type:text;comment:执行结果" json:"result"`                                                // 执行结果
	RunTimes      uint32     `gorm:"column:run_times;type:int unsigned;not null;comment:执行次数" json:"run_times"`                         // 执行次数
	MaxRetryTimes uint32     `gorm:"column:max_retry_times;type:int unsigned;not null;comment:最大重试次数" json:"max_retry_times"`           // 最大重试次数
	CreatedAt     *time.Time `gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:插入时间" json:"created_at"` // 插入时间
	UpdatedAt     *time.Time `gorm:"column:updated_at;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"` // 更新时间
}

// TableName Task's table name
func (*Task) TableName() string {
	return TableNameTask
}