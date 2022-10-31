package config

import (
	"flag"
	"fmt"
	logs "github.com/danbai225/go-logs"
	json "github.com/json-iterator/go"
	"os"
	"web_demo/pkg/env"
)

type Config struct {
	Env   string `json:"env"`
	Addr  string `json:"addr"`
	Mysql struct {
		User string `json:"user"`
		Pass string `json:"pass"`
		Addr string `json:"addr"`
		Name string `json:"name"`
	} `json:"mysql"`
	Redis struct {
		Pass string `json:"pass"`
		Addr string `json:"addr"`
		Db   int    `json:"db"`
	} `json:"redis"`
	AList struct {
		Url          string `json:"url"`
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
		WebdavUser   string `json:"webdav_user"`
		WebdavPass   string `json:"webdav_pass"`
	} `json:"a_list"`
	Dev struct {
		Username string `json:"username"`
	} `json:"dev"`
	Host string `json:"host"`
}

var config = new(Config)

func init() {
	cPath := flag.String("config", "", "手动指定配置文件路径")
	flag.Parse()
	if *cPath == "" {
		getenv := os.Getenv("web_demoConfig")
		if getenv != "" {
			*cPath = getenv
		}
	}
	if *cPath == "" {
		*cPath = fmt.Sprintf("config_%s.json", env.Active().Value())
	}
	data, err := os.ReadFile(*cPath)
	if err != nil {
		logs.Err(err, *cPath)
		os.Exit(1)
	}
	err = json.Unmarshal(data, config)
	if err != nil {
		logs.Err(err)
		os.Exit(1)
	}
	if env.Active().Value() != config.Env {
		logs.Warn("配置环境与参数环境不一致")
	}
}

func Get() Config {
	return *config
}
