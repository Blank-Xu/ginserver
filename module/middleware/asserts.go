package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Asserts(root string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != http.MethodGet && c.Request.Method != http.MethodHead {
			return
		}
		if !strings.HasPrefix(c.Request.URL.Path, root) {
			c.Next()
			return
		}
		// not done
		c.AbortWithStatus(http.StatusOK)
	}
}
