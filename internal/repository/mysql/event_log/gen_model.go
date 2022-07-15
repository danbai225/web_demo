package event_log

import "time"

// EventLog
//go:generate gormgen -structs EventLog -input .
type EventLog struct {
	Id        int64     // 自增id
	Type      string    // 操作类型
	Device    string    // 操作设备
	Username  string    // 用户名
	UserId    int64     // 用户id
	Ip        string    // ip地址
	Location  string    // 位置
	Content   string    // 具体内容
	CreatedAt time.Time `gorm:"time"` // 创建时间
}
