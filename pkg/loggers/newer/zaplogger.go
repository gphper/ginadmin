/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-07-10 09:59:02
 */
package newer

import (
	"io"
	"time"

	"github.com/gphper/ginadmin/configs"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewZapLogger(path string) *zap.Logger {
	// 设置一些基本日志格式 具体含义还比较好理解，直接看zap源码也不难懂
	encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder,
		TimeKey:     "ts",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		CallerKey:    "file",
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})
	// 获取 info、error日志文件的io.Writer 抽象 getWriter() 在下方实现
	infoWriter := getWriter(configs.RootPath + "/logs/%Y%m%d/" + path + "/demo_info.log")
	errorWriter := getWriter(configs.RootPath + "/logs/%Y%m%d/" + path + "/demo_error.log")

	// 最后创建具体的Logger
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), zapcore.InfoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(errorWriter), zapcore.ErrorLevel),
	)
	return zap.New(core, zap.AddCaller())
}

func getWriter(filename string) io.Writer {
	hook, err := rotatelogs.New(
		filename,
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour),
	)
	if err != nil {
		panic(err)
	}
	return hook
}
