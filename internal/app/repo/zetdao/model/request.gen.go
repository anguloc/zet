// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameRequest = "request"

// Request mapped from table <request>
type Request struct {
	ID          int64      `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true;comment:id" json:"id"`                          // id
	URL         string     `gorm:"column:url;type:varchar(1024);not null;comment:资源地址" json:"url"`                                    // 资源地址
	Mark        string     `gorm:"column:mark;type:varchar(64);not null;index:idx_mark,priority:1;comment:标识" json:"mark"`            // 标识
	RequestNum  int32      `gorm:"column:request_num;type:int;not null;comment:请求次数" json:"request_num"`                              // 请求次数
	Content     *string    `gorm:"column:content;type:longtext;comment:请求原始结果" json:"content"`                                        // 请求原始结果
	ParseStatus int32      `gorm:"column:parse_status;type:tinyint;not null;comment:请求内容解析状态" json:"parse_status"`                    // 请求内容解析状态
	CreatedAt   *time.Time `gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:插入时间" json:"created_at"` // 插入时间
	UpdatedAt   *time.Time `gorm:"column:updated_at;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"` // 更新时间
}

// TableName Request's table name
func (*Request) TableName() string {
	return TableNameRequest
}