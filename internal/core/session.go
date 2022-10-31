package core

import "web_demo/internal/data_obj/base"

type SessionUserInfo struct {
	*base.HttpBaseInfo
	Username string
	Email    string
}
