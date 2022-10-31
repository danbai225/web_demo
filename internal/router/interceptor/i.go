package interceptor

import (
	logs "github.com/danbai225/go-logs"
	"github.com/oschwald/geoip2-golang"
)

var ipDB *geoip2.Reader

func init() {
	var err error
	ipDB, err = geoip2.Open("assets/GeoLite2-City.mmdb")
	if err != nil {
		logs.Err(err.Error())
	}
}
