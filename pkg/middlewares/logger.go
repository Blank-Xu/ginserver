package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

// Logger middleware for gin
func Logger(logger *zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		status := c.Writer.Status()
		log := logger.With().
			Int("status", status).
			Str("method", c.Request.Method).
			Str("ip", c.ClientIP()).
			Str("path", c.Request.URL.Path).
			Str("raw_query", c.Request.URL.RawQuery).
			Str("referer", c.Request.Referer()).
			Str("user_agent", c.Request.UserAgent()).
			Float64("latency", float64(time.Now().Sub(start).Nanoseconds())/1000000.0).
			Logger()

		msg := "request"
		if len(c.Errors) > 0 {
			msg = c.Errors.String()
		}

		switch {
		case c.Writer.Status() >= http.StatusBadRequest && c.Writer.Status() < http.StatusInternalServerError:
			log.Warn().Msg(msg)
		case c.Writer.Status() >= http.StatusInternalServerError:
			log.Error().Msg(msg)
		default:
			log.Info().Msg(msg)
		}
	}
}
