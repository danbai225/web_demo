package webdav

import (
	"web_demo/internal/model/base"
	"web_demo/internal/utils"
)

type WebDav struct {
	ID             int64  `json:"id" gorm:"primaryKey;comment:id"`
	UserId         int64  `json:"user_id" gorm:"unique;comment:用户关联ID"`
	Url            string `json:"url" gorm:"type:varchar(256);comment:webdav地址"`
	Username       string `json:"username" gorm:"varchar(256):varchar;comment:webdav用户名"`
	Password       string `json:"-" gorm:"type:varchar(256);comment:webdav密码"`
	Public         bool   `json:"public" gorm:"comment:是否公开"`
	InternalServer bool   `json:"internal_server" gorm:"comment:使用内部服务器"`
	base.Base
}

func (w *WebDav) GetPass() string {
	return utils.Decrypt(w.Password)
}
func (w *WebDav) SetPass(pass string) error {
	encryption, err := utils.Encryption(pass)
	if err != nil {
		return err
	}
	w.Password = encryption
	return nil
}
