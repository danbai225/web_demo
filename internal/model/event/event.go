package event

import (
	"time"
)

type Event struct {
	ID        int64      `json:"id" gorm:"primaryKey;comment:id"`
	UserId    int64      `json:"user_id" gorm:"comment:用户id"`
	Type      string     `json:"type" gorm:"type:varchar(32);comment:类型"`
	Content   string     `json:"Content" gorm:"type:text;comment:内容"`
	Data      string     `json:"data" gorm:"type:longtext;comment:数据"`
	CreatedAt *time.Time `json:"created_at" gorm:"comment:创建时间"`
}
