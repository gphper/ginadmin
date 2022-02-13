/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-08-22 12:19:59
 */
package facade

import (
	"github.com/gphper/ginadmin/configs"
)

type Log interface {
	Info(string, map[string]string)
	Error(string, map[string]string)
}

func NewLogger(path string) (logger Log) {

	var logType string = configs.App.LogMedia

	switch logType {

	case "redis":
		logger = newRedisLog(path)

	default:
		logger = newZaplog(path)
	}

	return
}
