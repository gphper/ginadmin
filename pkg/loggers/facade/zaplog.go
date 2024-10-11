/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-08-22 12:22:19
 */
package facade

import (
	"context"

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

func (zlog ZapLog) Info(ctx context.Context, msg string, info map[string]string) {

	var fieldNum int

	value := ctx.Value("requestId")
	if value != nil {
		info["request_id"] = value.(string)
	}

	zapSlice := make([]zap.Field, len(info))

	for k, v := range info {
		zapSlice[fieldNum] = zap.String(k, v)
		fieldNum++
	}

	zlog.logger.Info(msg, zapSlice...)
}

func (zlog ZapLog) Error(ctx context.Context, msg string, info map[string]string) {

	var fieldNum int

	value := ctx.Value("requestId")
	if value != nil {
		info["request_id"] = value.(string)
	}

	zapSlice := make([]zap.Field, len(info))

	for k, v := range info {
		zapSlice[fieldNum] = zap.String(k, v)
		fieldNum++
	}
	zlog.logger.Error(msg, zapSlice...)
}
