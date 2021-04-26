package cron

import (
	"github.com/robfig/cron"
)

func init() {
	crontab := cron.New()
	//crontab.AddFunc("* * * * * *", myfunc)
	crontab.Start()
}

//func myfunc(){
//	fmt.Println("hello world")
//}
