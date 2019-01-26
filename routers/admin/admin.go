package admin

import (
	"github.com/gin-gonic/gin"

	"ginserver/controllers/admin"
)

func registerAdmin(r *gin.RouterGroup) {
	r.GET("admins/:id", admin.GetAdminById)
	r.GET("admins", admin.GetAdmin)
	r.POST("admins", admin.PostAdmin)
	r.PUT("admins/:id", admin.PutAdmin)
	r.DELETE("admins/:id", admin.DeleteAdmin)
}
