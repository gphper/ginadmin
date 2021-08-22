/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-08-22 12:22:29
 */
package facade

import (
	loggers "github/gphper/ginadmin/pkg/loggers"
)

type RedisLog struct {
	logger *loggers.RedisLogger
}

func NewRedisLog(path string) *RedisLog {
	return &RedisLog{
		logger: loggers.NewRedisLogger(path),
	}
}

func (rlog RedisLog) Info(msg string, info map[string]string) {

	rlog.logger.Info(msg, info)
}

func (rlog RedisLog) Error(msg string, info map[string]string) {
	rlog.logger.Error(msg, info)
}
