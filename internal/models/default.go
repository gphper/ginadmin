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

var mapDB map[string]*gorm.DB

type GaTabler interface {
	schema.Tabler
	FillData(*gorm.DB)
	GetConnName() string
}

type BaseModle struct {
	ConnName string `gorm:"-"`
}

func (b *BaseModle) TableName() string {
	return ""
}

func (b *BaseModle) FillData(db *gorm.DB) {}

func (b *BaseModle) GetConnName() string {
	return b.ConnName
}

//获取链接
func GetDB(model GaTabler) *gorm.DB {

	db, ok := mapDB[model.GetConnName()]
	if !ok {
		errMsg := fmt.Sprintf("connection name%s no exists", model.GetConnName())
		loggers.LogError("get_db_error", "GetDB", map[string]string{
			"msg": errMsg,
		})
	}
	return db
}

func Init() {
	mapDB = make(map[string]*gorm.DB)

	for _, mysqlConfig := range configs.App.Mysqls {
		var err error
		dns := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlConfig.UserName, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Database)
		db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		sqlDb, _ := db.DB()
		sqlDb.SetMaxOpenConns(mysqlConfig.MaxOpenConn)
		sqlDb.SetMaxIdleConns(mysqlConfig.MaxIdleConn)

		//注册回调函数
		RegisterCallback(db)

		mapDB[mysqlConfig.Name] = db
	}

	modelss := GetModels()
	if configs.App.Base.MigrateTable {

		for _, v := range modelss {
			db := GetDB(v.(GaTabler))
			err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(v)
			if err != nil {
				os.Exit(0)
			}
		}

	}

	if configs.App.Base.FillData {
		for _, v := range modelss {
			tabler := v.(GaTabler)
			db := GetDB(tabler)
			tabler.FillData(db)
		}
	}
}

func GetModels() []interface{} {
	return []interface{}{
		&AdminUsers{},
		&Article{},
		&UploadType{},
		&User{},
	}
}

func RegisterCallback(db *gorm.DB) {
	//注册创建数据回调
	db.Callback().Create().After("gorm:create").Register("my_plugin:after_create", func(db *gorm.DB) {
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
	db.Callback().Delete().After("gorm:delete").Register("my_plugin:after_delete", func(db *gorm.DB) {
		str := db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)
		loggers.LogInfo("sql", "delete sql", map[string]string{
			"info": str,
		})
	})
	//TODO 注册更新数据回调
	db.Callback().Update().After("gorm:update").Register("my_plugin:after_update", func(db *gorm.DB) {
		str := db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)
		loggers.LogInfo("sql", "update sql", map[string]string{
			"info": str,
		})
	})
}
