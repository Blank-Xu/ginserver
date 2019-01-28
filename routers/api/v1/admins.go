package v1

import (
	"github.com/gin-gonic/gin"

	"ginserver/controllers/api/v1/admin/admins"
)

func registerAdmin(r *gin.RouterGroup) {
	r.GET("admins/:id", admins.GetAdminById)
	r.GET("admins", admins.GetAdmin)
	r.POST("admins", admins.PostAdmin)
	r.PUT("admins/:id", admins.PutAdmin)
	r.DELETE("admins/:id", admins.DeleteAdmin)
}
