/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-10-19 19:51:15
 */
package cron

import (
	"github.com/robfig/cron"
)

func init() {
	crontab := cron.New()
	// crontab.AddFunc("0 0 23 * *", WriteLog)
	crontab.Start()
}

//func myfunc(){
//	fmt.Println("hello world")
//}
