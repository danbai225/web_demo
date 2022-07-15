package interceptor

import (
	"fmt"
	"net"

	"web_demo/internal/pkg/core"
)

type BaseInfo struct {
	Ip       string `json:"ip"`
	Device   string `json:"device"`
	Location string `json:"location"`
}

func (i *interceptor) BaseInfo() core.HandlerFunc {
	return func(c core.Context) {
		base := new(BaseInfo)
		if ip := c.GetHeader("X-Real-IP"); ip == "" {
			base.Ip = c.GetGin().ClientIP()
		} else {
			base.Ip = ip
		}
		base.Device = c.GetHeader("User-Agent")
		record, _ := i.ipDB.City(net.ParseIP(base.Ip))
		subdivisionName := ""
		if len(record.Subdivisions) > 0 {
			subdivisionName = record.Subdivisions[0].Names["zh-CN"]
		}
		base.Location = fmt.Sprintf("%s%s%s", record.Country.Names["zh-CN"], subdivisionName, record.City.Names["zh-CN"])
		c.SetVal("baseInfo", base)
	}
}
