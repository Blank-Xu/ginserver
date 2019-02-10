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
		req    = new(models.SUserLogin)
		newCtx = NewContext(ctx)
		err    error
	)
	if err = newCtx.ShouldBind(req); err != nil {
		newCtx.RespErrInvalidParams(err)
		return
	}
	// check user
	recordUser := &models.SUser{Username: req.Username}
	var has bool
	if has, err = recordUser.SelectOne(recordUser); err != nil {
		newCtx.RespErrDBError(err)
		return
	}
	if !has ||
		recordUser.Password != utils.GenPassword(req.Password, recordUser.Salt) {
		newCtx.RespErrInvalidParams()
		return
	}
	if recordUser.State == false {
		newCtx.RespErrForbidden()
		return
	}
	// check role
	recordRole := new(models.SRole)
	if has, err = recordRole.SelectOneByUserId(recordUser.Id); err != nil {
		newCtx.RespErrDBError(err)
		return
	}
	if !has || recordRole.State == false {
		newCtx.RespErrForbidden()
		return
	}
	// check success
	if err = newCtx.SessionCreate(recordUser.Id, recordRole.Id); err != nil {
		newCtx.RespErrInternalServerError(err)
		return
	}
	// cache user
	SetCacheUser(recordUser)
	newCtx.Log(models.LogTypeLogin, models.LogLevelInfo)
	newCtx.RespRedirect302("/admin")
}
