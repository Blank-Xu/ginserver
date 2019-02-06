package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"ginserver/models"
	"ginserver/modules/utils"
)

type login struct{}

func (p *login) registerRouter(r *gin.RouterGroup) {
	r.GET("login", p.get)
	r.POST("login", p.post)
}

func (p *login) get(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html",
		map[string]string{
			"Title":    "ginserverLogin",
			"UserName": "adming",
			"Password": "123456",
		})
}

func (p *login) post(ctx *gin.Context) {
	var (
		req    = new(models.SAdminLogin)
		newCtx = NewContext(ctx)
		err    error
	)
	if err = newCtx.Bind(req); err != nil {
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
	newCtx.RespRedirect302("/")
}
