/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-09-16 19:55:23
 */
package cron

import (
	"github.com/robfig/cron"
)

func init() {
	crontab := cron.New()
	//凌晨一点执行日志落盘
	crontab.AddFunc("0 0 1 * * *", WriteLog)
	crontab.Start()
}

//func myfunc(){
//	fmt.Println("hello world")
//}
