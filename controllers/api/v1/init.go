package v1

import (
	"ginserver/controllers/api/v1/admins"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.RouterGroup) {
	v1Group := r.Group("v1")
	new(admins.Admins).RegisterRouter(v1Group)
}
