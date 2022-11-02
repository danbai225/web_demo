package base

import (
	"sync"
)

type RegServer interface {
	CheckSever() error
	OverServer() error
}
type ServerManage struct {
	sMap sync.Map
}

func NewServerManage() *ServerManage {
	return &ServerManage{sMap: sync.Map{}}
}
func (m *ServerManage) Get(name string) interface{} {
	value, _ := m.sMap.Load(name)
	return value
}
func (m *ServerManage) Set(name string, sever RegServer) {
	m.sMap.Store(name, sever)
}
func (m *ServerManage) Check() error {
	var err error
	m.sMap.Range(func(key, value any) bool {
		err = value.(RegServer).CheckSever()
		if err != nil {
			return false
		}
		return true
	})
	return err
}
func (m *ServerManage) Over() error {
	var err error
	m.sMap.Range(func(key, value any) bool {
		err = value.(RegServer).OverServer()
		if err != nil {
			return false
		}
		return true
	})
	return err
}
