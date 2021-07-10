package models

import (
	"fmt"
	"ginadmin/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Db *gorm.DB

type GaTabler interface {
	schema.Tabler
	FillData()
}

type BaseModle struct {
}

func init() {
	var err error
	dns := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.App.MysqlConf.UserName, conf.App.MysqlConf.Password, conf.App.MysqlConf.Host, conf.App.MysqlConf.Database)
	Db, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	sqlDb, _ := Db.DB()
	sqlDb.SetMaxOpenConns(conf.App.MaxOpenConn)
	sqlDb.SetMaxIdleConns(conf.App.MaxIdleConn)

	//注册回调函数
	RegisterCallback()
}

func GetModels() []interface{} {
	return []interface{}{
		&AdminUsers{},
		&AdminGroup{},
	}
}

func RegisterCallback() {
	//注册创建数据回调
	Db.Callback().Create().After("gorm:create").Register("my_plugin:after_create", func(db *gorm.DB) {
		str := fmt.Sprintf("sql语句：%s 参数：%s", db.Statement.SQL.String(), db.Statement.Vars)
		fmt.Println(str)
	})
	// Db.Callback().Query().After("gorm:query").Register("my_plugin:after_select", func(db *gorm.DB) {
	// 	str := fmt.Sprintf("sql语句：%s 参数：%s", db.Statement.SQL.String(), db.Statement.Vars)
	// 	fmt.Println(str)
	// })
	//TODO 注册删除数据回调

	//TODO 注册更新数据回调
}
