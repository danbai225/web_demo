package file

import "time"

type File struct {
	ID        int64      `json:"id" gorm:"primaryKey;comment:id"`
	UserId    int64      `json:"user_id" gorm:"index;comment:用户id"`
	Type      string     `json:"type" gorm:"type:varchar(32);comment:类型"`
	Filename  string     `json:"filename" gorm:"type:varchar(256);comment:文件名"`
	Path      string     `json:"path" gorm:"type:varchar(256);comment:路径"`
	Sha256    string     `json:"sha256" gorm:"type:varchar(256);comment:Sha256"`
	CreatedAt *time.Time `json:"created_at" gorm:"comment:创建时间"`
}
