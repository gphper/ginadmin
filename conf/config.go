package conf

import (
	"fmt"
	"ginadmin/comment"
	"gopkg.in/ini.v1"
)

type AppConf struct {
	MysqlConf `ini:"mysql"`
	SessionConf `ini:"session"`
	BaseConf `ini:"base"`
}

type MysqlConf struct {
	Host string `ini:"host"`
	Port string `ini:"port"`
	UserName string `ini:"username"`
	Password string `ini:"password"`
	Database string `ini:"database"`
}

type SessionConf struct {
	SessionName string `ini:"session_name"`
}

type BaseConf struct {
	Port string `ini:"port"`
}

var App = new(AppConf)

//初始化配置文件
func init(){
	path,err := comment.RootPath()
	if err != nil {
		fmt.Printf("get root path err:%v",err)
	}
	fmt.Println(path)
	err = ini.MapTo(App,path+"/conf/config.ini")
	if err != nil{
		fmt.Printf("load ini err:%v",err)
	}
}