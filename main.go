package main

import (
	"context"
	"flag"
	logs "github.com/danbai225/go-logs"
	"net/http"
	"time"
	"web_demo/config"
	"web_demo/internal/core"
	"web_demo/internal/model"
	"web_demo/internal/repository/mysql"
	"web_demo/internal/repository/redis"
	"web_demo/internal/router"
	"web_demo/internal/services"
	"web_demo/pkg/shutdown"
)

func main() {
	c := config.Get()
	//http://patorjk.com/software/taag/#p=display&h=0&f=Ogre&t=web_demo
	logs.Info("env:", c.Env, "                 _                 _                         \n__      __  ___ | |__           __| |  ___  _ __ ___    ___  \n\\ \\ /\\ / / / _ \\| '_ \\         / _` | / _ \\| '_ ` _ \\  / _ \\ \n \\ V  V / |  __/| |_) |       | (_| ||  __/| | | | | || (_) |\n  \\_/\\_/   \\___||_.__/  _____  \\__,_| \\___||_| |_| |_| \\___/ \n                       |_____|                               ")
	//mysql
	db, err := mysql.New(c.Mysql.User, c.Mysql.Pass, c.Mysql.Addr, c.Mysql.Name)
	if err != nil {
		logs.Err(err)
		return
	}
	if flag.Arg(0) == "syncDB" {
		logs.Info("数据库初始化")
		err = db.GetDb().AutoMigrate(model.Export()...)
		if err != nil {
			logs.Err(err)
		} else {
			logs.Info("初始化完成")
		}
		return
	}
	//redis
	rdb, err := redis.New(c.Redis.Addr, c.Redis.Pass, c.Redis.Db)
	if err != nil {
		logs.Err(err)
		return
	}
	services.InitService(db, rdb)
	//http服务
	server := &http.Server{
		Addr:    c.Addr,
		Handler: core.New(router.RegRouter, db, rdb),
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logs.Err("http services startup err", err)
		}
	}()
	// 优雅关闭
	shutdown.NewHook().Close(
		// 关闭 http services
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()
			if err := server.Shutdown(ctx); err != nil {
				logs.Err("services shutdown err", err)
			}
		},
		//service
		func() {
			err := services.Close()
			if err != nil {
				logs.Err(err)
				return
			}
		},
		// 关闭 mysql services
		func() {
			if err := db.DbClose(); err != nil {
				logs.Err("mysql shutdown err", err)
			}
		},
		// 关闭 redis services
		func() {
			if err := rdb.Close(); err != nil {
				logs.Err("redis shutdown err", err)
			}
		},
	)
}
