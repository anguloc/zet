// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package zet_model

const TableNameRequest = "request"

// Request mapped from table <request>
type Request struct {
	ID          int64   `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true;comment:id" json:"id"`
	URL         string  `gorm:"column:url;type:varchar(1024);not null;comment:资源地址" json:"url"`
	RequestNum  int32   `gorm:"column:request_num;type:int;not null;comment:请求次数" json:"request_num"`
	Content     *string `gorm:"column:content;type:text;comment:请求原始结果" json:"content"`
	AddTime     int32   `gorm:"column:add_time;type:int;not null" json:"add_time"`
	UpdateTime  int32   `gorm:"column:update_time;type:int;not null" json:"update_time"`
	DeletedTime int32   `gorm:"column:deleted_time;type:int;not null" json:"deleted_time"`
}

// TableName Request's table name
func (*Request) TableName() string {
	return TableNameRequest
}
