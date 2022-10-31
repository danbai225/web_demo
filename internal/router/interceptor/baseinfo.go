package interceptor

import (
	"fmt"
	"net"
	"web_demo/internal/core"
	"web_demo/internal/data_obj/base"
)

func BaseInfo() {
	baseInfo := new(base.HttpBaseInfo)
	c := core.GetContext()
	if ip := c.GetHeader("X-Real-IP"); ip == "" {
		baseInfo.Ip = c.GetGin().ClientIP()
	} else {
		baseInfo.Ip = ip
	}
	baseInfo.Device = c.GetHeader("User-Agent")
	record, _ := ipDB.City(net.ParseIP(baseInfo.Ip))
	subdivisionName := ""
	if len(record.Subdivisions) > 0 {
		subdivisionName = record.Subdivisions[0].Names["zh-CN"]
	}
	baseInfo.Location = fmt.Sprintf("%s%s%s", record.Country.Names["zh-CN"], subdivisionName, record.City.Names["zh-CN"])
	c.SetSessionUserInfo(&core.SessionUserInfo{HttpBaseInfo: baseInfo})
}
