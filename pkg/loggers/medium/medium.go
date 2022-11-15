/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-08-22 12:49:26
 */
package medium

import (
	"context"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
	"time"

	"github.com/gphper/ginadmin/pkg/loggers/facade"

	"github.com/gin-gonic/gin"
)

func GinLog(logger facade.Log, timeFormat string, utc bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		if utc {
			end = end.UTC()
		}

		ctx, _ := c.Get("ctx")

		if len(c.Errors) > 0 {
			infoMap := make(map[string]string, len(c.Errors))
			for ek, e := range c.Errors.Errors() {
				infoMap[strconv.Itoa(ek)] = e
			}
			logger.Error(ctx.(context.Context), "error msg", infoMap)
		} else {

			logger.Info(ctx.(context.Context), path, map[string]string{
				"status":     strconv.Itoa(c.Writer.Status()),
				"method":     c.Request.Method,
				"path":       path,
				"query":      query,
				"ip":         c.ClientIP(),
				"user-agent": c.Request.UserAgent(),
				"time":       end.Format(timeFormat),
				"latency":    latency.String(),
			})

		}
	}
}

func RecoveryWithLog(logger facade.Log, stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, _ := c.Get("ctx")

		defer func() {
			if err := recover(); err != nil {

				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {

					logger.Error(ctx.(context.Context), c.Request.URL.Path, map[string]string{
						"error":   err.(string),
						"request": string(httpRequest),
					})

					c.Error(err.(error))
					c.Abort()
					return
				}

				if stack {
					logger.Error(ctx.(context.Context), "[Recovery from panic]", map[string]string{
						"time":    time.Now().String(),
						"error":   err.(string),
						"request": string(httpRequest),
						"stack":   string(debug.Stack()),
					})
				} else {
					logger.Error(ctx.(context.Context), "[Recovery from panic]", map[string]string{
						"time":    time.Now().String(),
						"error":   err.(string),
						"request": string(httpRequest),
					})
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
