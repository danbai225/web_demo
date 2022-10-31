package base

type RegServer interface {
	CheckSever() error
	OverServer() error
}
type ServerManage struct {
	sMap map[string]RegServer
}

func NewServerManage() *ServerManage {
	return &ServerManage{sMap: make(map[string]RegServer)}
}
func (m *ServerManage) Get(name string) interface{} {
	return m.sMap[name]
}
func (m *ServerManage) Set(name string, sever RegServer) {
	m.sMap[name] = sever
}
func (m *ServerManage) Check() error {
	for _, v := range m.sMap {
		err := v.CheckSever()
		if err != nil {
			return err
		}
	}
	return nil
}
func (m *ServerManage) Over() error {
	for _, v := range m.sMap {
		err := v.OverServer()
		if err != nil {
			return err
		}
	}
	return nil
}
