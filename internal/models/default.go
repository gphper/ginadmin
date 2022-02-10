/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-07-08 20:12:04
 */
package models

import (
	"fmt"
	"os"

	"github.com/gphper/ginadmin/configs"
	"github.com/gphper/ginadmin/pkg/loggers"

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
	dns := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", configs.App.MysqlConf.UserName, configs.App.MysqlConf.Password, configs.App.MysqlConf.Host, configs.App.MysqlConf.Port, configs.App.MysqlConf.Database)
	Db, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDb, _ := Db.DB()
	sqlDb.SetMaxOpenConns(configs.App.MaxOpenConn)
	sqlDb.SetMaxIdleConns(configs.App.MaxIdleConn)

	modelss := GetModels()
	if configs.App.MigrateTable {
		err := Db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(modelss...)
		if err != nil {
			os.Exit(0)
		}
	}

	if configs.App.FillData {
		for _, v := range modelss {
			tabler := v.(GaTabler)
			tabler.FillData()
		}
	}

	//注册回调函数
	RegisterCallback()
}

func GetModels() []interface{} {
	return []interface{}{
		&AdminUsers{},
		&AdminGroup{},
		&Article{},
		&UploadType{},
		&User{},
	}
}

func RegisterCallback() {
	//注册创建数据回调
	Db.Callback().Create().After("gorm:create").Register("my_plugin:after_create", func(db *gorm.DB) {
		str := db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)
		loggers.LogInfo("sql", "create sql", map[string]string{
			"info": str,
		})
	})
	// Db.Callback().Query().After("gorm:query").Register("my_plugin:after_select", func(db *gorm.DB) {
	// 	str := fmt.Sprintf("sql语句：%s 参数：%s", db.Statement.SQL.String(), db.Statement.Vars)
	// 	fmt.Println(str)
	// })
	//TODO 注册删除数据回调
	Db.Callback().Delete().After("gorm:delete").Register("my_plugin:after_delete", func(db *gorm.DB) {
		str := db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)
		loggers.LogInfo("sql", "delete sql", map[string]string{
			"info": str,
		})
	})
	//TODO 注册更新数据回调
	Db.Callback().Update().After("gorm:update").Register("my_plugin:after_update", func(db *gorm.DB) {
		str := db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)
		loggers.LogInfo("sql", "update sql", map[string]string{
			"info": str,
		})
	})
}
