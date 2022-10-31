package user

import "web_demo/internal/model/base"

type User struct {
	ID           int64  `json:"id" gorm:"primaryKey;comment:id"`
	Username     string `json:"username" gorm:"type:varchar(256);unique;comment:用户名"`
	Email        string `json:"email" gorm:"type:varchar(256);unique;comment:邮箱"`
	Avatar       string `json:"avatar" gorm:"type:text;comment:头像"`
	Authenticate string `json:"-" gorm:"type:varchar(256);comment:认证码"`
	base.Base
}
