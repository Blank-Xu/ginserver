package middleware

import (
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// logrus middleware for gin
func Logger(log *logrus.Logger, assetsFile string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if gin.Mode() != gin.DebugMode && strings.HasPrefix(c.Request.URL.Path, assetsFile) {
			return
		}
		start := time.Now()
		c.Next()
		status := c.Writer.Status()
		entry := log.WithFields(logrus.Fields{
			"method": c.Request.Method,
			//"host":       c.Request.Host,
			"ip":         c.ClientIP(),
			"path":       c.Request.URL.Path,
			"referer":    c.Request.Referer(),
			"user_agent": c.Request.UserAgent(),
			"status":     status,
		})

		if gin.Mode() == gin.DebugMode {
			entry = entry.WithField("latency", strconv.FormatFloat(float64(
				time.Now().Sub(start).Nanoseconds())/1000000.0, 'f', -1, 32)+" ms")
		}

		switch {
		case status > 499:
			entry.Error(c.Errors.String())
		case status > 399:
			entry.Warn(c.Errors.String())
		default:
			entry.Info(c.Errors.String())
		}
	}
}
