/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-09-14 20:21:40
 */
package cron

import (
	"bufio"
	"fmt"
	globalRedis "github/gphper/ginadmin/internal/redis"
	"github/gphper/ginadmin/pkg/comment"
	"io"
	"strings"
	"sync"
)

/**
* 创建目录
**/
func WriteLog() {

	var wg sync.WaitGroup

	// date := time.Now().Local().Format("20060102")
	date := "20211201"

	pattern := "logs:" + date + ":*"
	keys, _ := globalRedis.RedisClient.Keys(pattern).Result()

	rootPath, _ := comment.RootPath()

	for _, key := range keys {
		path := strings.ReplaceAll(key, ":", "/")
		file, err := comment.OpenFile(rootPath + "/" + path + ".log")
		if err == nil {
			wg.Add(1)
			go writeFile(key, file, &wg)
		}
	}

	wg.Wait()
}

/**
* 内容写入到文件中
**/
func writeFile(key string, file io.Writer, wg *sync.WaitGroup) {
	defer wg.Done()

	var start int64 = 0
	var end int64 = 2
	for {
		// end := start + page*limit
		logs, _ := globalRedis.RedisClient.LRange(key, start, end).Result()
		fmt.Println(logs)
		if len(logs) > 0 {
			w := bufio.NewWriter(file)
			for _, log := range logs {
				_, writeerr := fmt.Fprintln(w, log)
				if writeerr != nil {
					continue
				}
			}
			finishErr := w.Flush()
			if finishErr == nil {
				//删除redis中已写完数据
				globalRedis.RedisClient.LTrim(key, end+1, -1).Result()
			}
		} else {
			break
		}
	}
}
