package base

type ReturnMsg struct {
	Code  int64       `json:"code"`  //业务状态码
	Msg   string      `json:"msg"`   //消息
	Data  interface{} `json:"data"`  //数据
	Trace string      `json:"trace"` //请求ID
}
type GlobalIdStructure struct {
	Id int64 `json:"id" form:"id"`
}
type HttpBaseInfo struct {
	Device   string `json:"device"`
	Ip       string `json:"ip"`
	Location string `json:"location"`
}
