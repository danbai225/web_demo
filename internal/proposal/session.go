package proposal

import "encoding/json"

// SessionUserInfo 当前用户会话信息
type SessionUserInfo struct {
	UserType int64  `json:"user_type"`
	Device   string `json:"device"`   //设备名称
	IP       string `json:"ip"`       //ip地址
	Location string `json:"location"` //位置
}

// Marshal 序列化到JSON
func (user *SessionUserInfo) Marshal() (jsonRaw []byte) {
	jsonRaw, _ = json.Marshal(user)
	return
}
