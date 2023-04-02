/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-10-19 19:51:15
 */
package cron

import (
	"context"
	"fmt"
	"runtime"

	"github.com/gphper/ginadmin/pkg/loggers"
	"github.com/robfig/cron/v3"
)

var cronHandle *cron.Cron

func Init() {
	cronHandle = cron.New(cron.WithSeconds(), cron.WithChain(cronRecover()))
	// cronHandle.AddFunc("0 0 23 * * ?", WriteLog)

	// cronHandle.AddFunc("* * * * * *", func() {
	// 	fmt.Println("hello world")
	// 	panic("hello panic")
	// })

	cronHandle.Start()
}

func GraceClose() (bool, context.Context) {
	if cronHandle == nil {
		return false, context.Background()
	}
	return true, cronHandle.Stop()
}

func cronRecover() cron.JobWrapper {
	return func(j cron.Job) cron.Job {
		return cron.FuncJob(func() {
			defer func() {
				if r := recover(); r != nil {
					const size = 64 << 10
					buf := make([]byte, size)
					buf = buf[:runtime.Stack(buf, false)]
					err, ok := r.(error)
					if !ok {
						err = fmt.Errorf("%v", r)
					}
					loggers.LogError(context.Background(), "cron", err.Error(), map[string]string{
						"detail": string(buf),
					})
				}
			}()
			j.Run()
		})
	}
}

//func myfunc(){
//	fmt.Println("hello world")
//}
