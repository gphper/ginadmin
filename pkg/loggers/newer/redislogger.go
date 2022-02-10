/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-08-22 15:03:08
 */

package newer

import (
	"encoding/json"
	"time"

	globalRedis "github.com/gphper/ginadmin/internal/redis"

	"github.com/go-redis/redis"
)

type RedisLogger struct {
	Client *redis.Client
	Path   string
}

func NewRedisLogger(path string) *RedisLogger {

	return &RedisLogger{
		Client: globalRedis.RedisClient,
		Path:   path,
	}
}

func (logger *RedisLogger) Info(msg string, info map[string]string) {
	info["level"] = "info"
	info["msg"] = msg
	info["ts"] = time.Now().String()

	str, _ := json.Marshal(info)
	time := time.Now().Format("20060102")

	logger.Client.LPush("logs:"+time+":"+logger.Path+":info", string(str))
}

func (logger *RedisLogger) Error(msg string, info map[string]string) {
	info["level"] = "error"
	info["msg"] = msg
	info["ts"] = time.Now().String()

	str, _ := json.Marshal(info)
	time := time.Now().Format("20060102")

	logger.Client.LPush("logs:"+time+":"+logger.Path+":error", string(str))
}
