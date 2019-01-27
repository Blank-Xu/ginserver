package v1

import "github.com/gin-gonic/gin"

func Init(r *gin.RouterGroup) {
	v1Router := r.Group("v1")
	registerAdmin(v1Router)
}
