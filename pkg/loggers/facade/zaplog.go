/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-08-22 12:22:19
 */
package facade

import (
	"github.com/gphper/ginadmin/pkg/loggers/newer"

	"go.uber.org/zap"
)

type ZapLog struct {
	logger *zap.Logger
}

func newZaplog(path string) *ZapLog {
	return &ZapLog{
		logger: newer.NewZapLogger(path),
	}
}

func (zlog ZapLog) Info(msg string, info map[string]string) {
	zapSlice := make([]zap.Field, len(info))
	var fieldNum int
	for k, v := range info {
		zapSlice[fieldNum] = zap.String(k, v)
		fieldNum++
	}
	zlog.logger.Info(msg, zapSlice...)
}

func (zlog ZapLog) Error(msg string, info map[string]string) {
	zapSlice := make([]zap.Field, len(info))
	var fieldNum int
	for k, v := range info {
		zapSlice[fieldNum] = zap.String(k, v)
		fieldNum++
	}
	zlog.logger.Error(msg, zapSlice...)
}
