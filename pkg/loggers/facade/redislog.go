/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-08-22 12:22:29
 */
package facade

import (
	"github.com/gphper/ginadmin/pkg/loggers/newer"
)

type RedisLog struct {
	logger *newer.RedisLogger
}

func newRedisLog(path string) *RedisLog {

	return &RedisLog{
		logger: newer.NewRedisLogger(path),
	}
}

func (rlog RedisLog) Info(msg string, info map[string]string) {

	rlog.logger.Info(msg, info)
}

func (rlog RedisLog) Error(msg string, info map[string]string) {
	rlog.logger.Error(msg, info)
}
