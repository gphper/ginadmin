/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-08-22 12:19:59
 */
package facade

import (
	"context"

	"github.com/gphper/ginadmin/configs"
)

type Log interface {
	Info(context.Context, string, map[string]string)
	Error(context.Context, string, map[string]string)
}

func NewLogger(path string) (logger Log) {

	var logType string = configs.App.Base.LogMedia

	switch logType {

	case "redis":
		logger = newRedisLog(path)

	default:
		logger = newZaplog(path)
	}

	return
}
