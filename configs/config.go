//go:build !embed

/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-07-04 11:58:45
 */

package configs

import (
	"flag"
	"io/ioutil"
	"testing"

	"github.com/gphper/ginadmin/pkg/utils/filesystem"

	"gopkg.in/yaml.v2"
)

var RootPath string

type AppConf struct {
	Mysqls  []MysqlConf `yaml:"mysql" json:"mysql"`
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

var App *AppConf

// 初始化配置文件
func Init(path string) error {

	var err error
	if path == "" {
		RootPath, err = filesystem.RootPath()
		if err != nil {

			return err
		}
	} else {
		RootPath = path
	}

	//否则执行 go test 报错
	testing.Init()
	flag.Parse()

	yamlFile, err := ioutil.ReadFile(RootPath + "/configs/config.yaml")
	if err != nil {

		return err
	}
	err = yaml.Unmarshal(yamlFile, &App)
	if err != nil {

		return err
	}

	return nil
}
