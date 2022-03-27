/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-09-14 20:21:40
 */
package cron

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"sync"
	"time"

	"github.com/gphper/ginadmin/configs"
	globalRedis "github.com/gphper/ginadmin/internal/redis"
	"github.com/gphper/ginadmin/pkg/utils/filesystem"
	gstrings "github.com/gphper/ginadmin/pkg/utils/strings"
)

/**
* 创建目录
**/
func WriteLog() {

	var wg sync.WaitGroup

	date := time.Now().AddDate(0, 0, -1).Local().Format("20060102")

	pattern := "logs:" + date + ":*"
	keys, _ := globalRedis.RedisClient.Keys(pattern).Result()

	for _, key := range keys {
		path := strings.ReplaceAll(key, ":", "/")

		file, err := filesystem.OpenFile(gstrings.JoinStr(configs.RootPath, "/", path, ".log"))
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
