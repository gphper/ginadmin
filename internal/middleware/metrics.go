/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-08-04 22:39:47
 */
package middleware

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	gapro "github.com/gphper/ginmonitor/prometheus"
	"github.com/prometheus/client_golang/prometheus"
)

func Metrics() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		//统计qps
		gapro.HttpRequestCounter.With(prometheus.Labels{"method": c.Request.Method, "path": c.FullPath()}).Inc()
		c.Next()
		endTime := time.Now()
		//统计相应时长
		gapro.HttpRquestTime.With(prometheus.Labels{"method": c.Request.Method, "path": c.FullPath()}).Set(float64(endTime.Sub(startTime)))
		//统计相应状态
		gapro.HttpRquestStatus.With(prometheus.Labels{"method": c.Request.Method, "path": c.FullPath(), "status": strconv.Itoa(c.Writer.Status())}).Inc()
	}
}
