package middlewares

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// logrus middleware for gin
func Logger(log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		status := c.Writer.Status()
		entry := log.WithFields(logrus.Fields{
			"method":     c.Request.Method,
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
