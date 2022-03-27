//+build !embed

/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-07-04 11:58:45
 */

package configs

import (
	"flag"
	"fmt"

	"github.com/gphper/ginadmin/pkg/utils/filesystem"

	"gopkg.in/ini.v1"
)

var RootPath string

type AppConf struct {
	MysqlConf   `ini:"mysql"`
	RedisConf   `ini:"redis"`
	SessionConf `ini:"session"`
	BaseConf    `ini:"base"`
}

type MysqlConf struct {
	Host        string `ini:"host"`
	Port        string `ini:"port"`
	UserName    string `ini:"username"`
	Password    string `ini:"password"`
	Database    string `ini:"database"`
	MaxOpenConn int    `ini:"max_open_conn"`
	MaxIdleConn int    `ini:"max_idle_conn"`
}

type RedisConf struct {
	Addr     string `ini:"addr"`
	Db       int    `ini:"db"`
	Password string `ini:"password"`
}

type SessionConf struct {
	SessionName string `ini:"session_name"`
}

type BaseConf struct {
	Port         string `ini:"port"`
	Host         string `ini:"host"`
	FillData     bool   `ini:"fill_data"`
	MigrateTable bool   `ini:"migrate_table"`
	LogMedia     string `ini:"log_media"`
}

var App = new(AppConf)

//初始化配置文件
func init() {

	path, err := filesystem.RootPath()
	if err != nil {
		fmt.Printf("get root path err:%v", err)
	}

	flag.StringVar(&RootPath, "root_path", path, "root path")

	flag.Parse()
	fmt.Println(RootPath)
	err = ini.MapTo(App, RootPath+"/configs/config.ini")
	if err != nil {
		fmt.Printf("load ini err:%v", err)
	}
}
