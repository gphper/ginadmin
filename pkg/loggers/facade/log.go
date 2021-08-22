/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-08-22 12:19:59
 */
package facade

type Log interface {
	Info(string,map[string]string)
	Error(string,map[string]string)
}
