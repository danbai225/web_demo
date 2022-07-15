package base

type ReturnMsg struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
type GlobalIdStructure struct {
	Id int64 `json:"id" form:"id"`
}
