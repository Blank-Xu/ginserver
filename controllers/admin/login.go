package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"ginserver/models"
	"ginserver/modules/utils"
)

type ControllerLogin struct{}

func (p *ControllerLogin) Get(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html",
		map[string]string{
			"Title":    "Login",
			"Username": "admin",
			"Password": "123456",
		})
}

func (p *ControllerLogin) Post(ctx *gin.Context) {
	var (
		req    = new(models.SAdminLogin)
		newCtx = NewContext(ctx)
		err    error
	)
	if err = newCtx.ShouldBind(req); err != nil {
		newCtx.RespErrInvalidParams(err)
		return
	}

	recordAdmin := &models.SAdmin{Username: req.Username}
	var has bool
	if has, err = recordAdmin.SelectOne(recordAdmin); err != nil {
		newCtx.RespErrDBError(err)
		return
	}
	if !has ||
		recordAdmin.Password != utils.GenPassword(req.Password, recordAdmin.Salt) {
		newCtx.RespErrInvalidParams()
		return
	}
	if recordAdmin.State == false {
		newCtx.RespErrForbidden()
		return
	}

	if err = newCtx.SessionCreate(recordAdmin.Id, recordAdmin.RoleId); err != nil {
		newCtx.RespErrInternalServerError(err)
		return
	}
	newCtx.Log(models.LogTypeLogin)
	newCtx.RespRedirect302("/admin")
}
