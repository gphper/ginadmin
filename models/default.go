package models

import (
	"fmt"
	"ginadmin/comment"
	"ginadmin/conf"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

func init() {
	var err error
	dns := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.App.MysqlConf.UserName, conf.App.MysqlConf.Password, conf.App.MysqlConf.Host, conf.App.MysqlConf.Database)
	Db, err = gorm.Open("mysql", dns)
	if err != nil {
		fmt.Println(err)
	}
	Db.DB().SetMaxOpenConns(conf.App.MaxOpenConn)
	Db.DB().SetMaxIdleConns(conf.App.MaxIdleConn)
	//自动注册数据表
	Db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&AdminUsers{}, &AdminGroup{})
	//注册回调函数
	RegisterCallback()
	//配置文件判断是否填充数据
	if conf.App.BaseConf.FillData {
		FillData()
	}
}

func FillData() {
	//填充管理用户组
	adminGroup := AdminGroup{
		GroupId:   1,
		GroupName: "管理员组",
		Privs:     "{\"all\":{}}",
	}
	Db.Save(&adminGroup)
	//初始化管理员
	salt := comment.RandString(6)
	passwordSalt := comment.Encryption("111111", salt)
	adminUser := AdminUsers{
		Uid:       1,
		GroupId:   1,
		Username:  "admin",
		Nickname:  "管理员",
		Password:  passwordSalt,
		Phone:     "",
		LastLogin: "",
		Salt:      salt,
		ApiToken:  "",
	}
	Db.Save(&adminUser)
}

func RegisterCallback() {
	//注册创建数据回调
	Db.Callback().Create().After("gorm:create").Register("my_plugin:after_create", func(scope *gorm.Scope) {
		str := fmt.Sprintf("sql语句：%s 参数：%s", scope.SQL, scope.SQLVars)
		fmt.Println(str)
	})
	//TODO 注册删除数据回调

	//TODO 注册更新数据回调
}
