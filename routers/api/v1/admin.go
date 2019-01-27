package v1

import (
	"github.com/gin-gonic/gin"

	"ginserver/controllers/api/v1/admin"
)

func registerAdmin(r *gin.RouterGroup) {
	r.GET("admins/:id", admin.GetAdminById)
	r.GET("admins", admin.GetAdmin)
	r.POST("admins", admin.PostAdmin)
	r.PUT("admins/:id", admin.PutAdmin)
	r.DELETE("admins/:id", admin.DeleteAdmin)
}
