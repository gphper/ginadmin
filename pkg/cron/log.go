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
	"time"
)

/**
* 创建目录
**/
func WriteLog() {

	date := time.Now().AddDate(0, 0, -2).Local().Format("20060102")

	pattern := "logs:" + date + ":*"
	keys, _ := globalRedis.RedisClient.Keys(pattern).Result()

	rootPath, _ := comment.RootPath()

	for _, key := range keys {
		path := strings.ReplaceAll(key, ":", "/")
		file, err := comment.OpenFile(rootPath + "/" + path + ".log")
		if err == nil {
			go writeFile(key, file)
		}
	}

}

/**
* 内容写入到文件中
**/
func writeFile(key string, file io.Writer) {
	var start int64 = 0
	var limit int64 = 100
	var page int64 = 1
	for {
		end := start + page*limit
		logs, _ := globalRedis.RedisClient.LRange(key, start, end).Result()
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
				globalRedis.RedisClient.LTrim(key, end, -1)
			}
		} else {
			break
		}

		page++
		start = end
	}
}
