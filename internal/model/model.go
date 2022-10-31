package model

import (
	"web_demo/internal/model/event"
	"web_demo/internal/model/file"
	"web_demo/internal/model/user"
	"web_demo/internal/model/webdav"
)

func Export() []interface{} {
	return []interface{}{&user.User{}, &webdav.WebDav{}, &event.Event{}, &file.File{}}
}
