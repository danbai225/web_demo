package base

type ServerManage struct {
	Map map[string]interface{}
}

func (ServerManage) New() *ServerManage {
	return &ServerManage{Map: make(map[string]interface{})}
}
func (m *ServerManage) Get(name string) interface{} {
	return m.Map[name]
}
func (m *ServerManage) Set(name string, sever interface{}) {
	m.Map[name] = sever
}
