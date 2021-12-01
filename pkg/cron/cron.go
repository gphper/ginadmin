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
	crontab.AddFunc("* * * * * *", WriteLog)
	crontab.Start()
}

//func myfunc(){
//	fmt.Println("hello world")
//}
