package test

import (
	"fmt"
	"net"
	"testing"

	"github.com/oschwald/geoip2-golang"
)

func TestIP(t *testing.T) {
	ipDB, err := geoip2.Open("../assets/GeoLite2-City.mmdb")
	if err != nil {
		t.Error(err)
	}
	record, _ := ipDB.City(net.ParseIP("118.117.51.31"))
	subdivisionName := ""
	if len(record.Subdivisions) > 0 {
		subdivisionName = record.Subdivisions[0].Names["zh-CN"]
	}
	Location := fmt.Sprintf("%s%s%s", record.Country.Names["zh-CN"], subdivisionName, record.City.Names["zh-CN"])
	println(Location)
}
