//go:build embed
// +build embed

/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-07-29 19:47:42
 */

package configs

import (
	"bytes"
	_ "embed"
	"fmt"

	"gopkg.in/yaml.v2"
)

type AppConf struct {
	Mysql   MysqlConf   `yaml:"mysql" json:"mysql"`
	Redis   RedisConf   `yaml:"redis" json:"redis"`
	Session SessionConf `yaml:"session" json:"session"`
	Base    BaseConf    `yaml:"base" json:"base"`
}

type MysqlConf struct {
	Name        string `yaml:"name" json:"name"`
	Host        string `yaml:"host" json:"host"`
	Port        string `yaml:"port" json:"port"`
	UserName    string `yaml:"username" json:"username"`
	Password    string `yaml:"password" json:"password"`
	Database    string `yaml:"database" json:"database"`
	MaxOpenConn int    `yaml:"max_open_conn" json:"max_open_conn"`
	MaxIdleConn int    `yaml:"max_idle_conn" json:"max_idle_conn"`
}

type RedisConf struct {
	Addr     string `yaml:"addr"`
	Db       int    `yaml:"db"`
	Password string `yaml:"password"`
}

type SessionConf struct {
	SessionName string `yaml:"session_name"`
}

type BaseConf struct {
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	LogMedia string `yaml:"log_media"`
}

var App = new(AppConf)

//go:embed config.yaml
var iniStr string

//初始化配置文件
func Init() {

	err = yaml.Unmarshal(bytes.NewReader([]byte(iniStr)), App)
	if err != nil {
		fmt.Println(err.Error())
	}
}
