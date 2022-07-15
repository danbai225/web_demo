package key_map

// KeyMap
//go:generate gormgen -structs KeyMap -input .
type KeyMap struct {
	Id  int64  //
	K   string // 主键
	Val string // 值
}
