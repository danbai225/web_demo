package event

import (
	logs "github.com/danbai225/go-logs"
)

func (s *service) New(Type, Content string, data interface{}) {
	//save to db
	logs.Println(Type, Content)
}
