package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"ginserver/models"
	"ginserver/modules/config"
	"ginserver/modules/log"
	"ginserver/modules/utils"
)

type ControllerLogin struct{}

func (p *ControllerLogin) Get(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html",
		map[string]string{
			"AppName":  config.GetConfig().AppName,
			"Title":    "Login",
			"Username": "admin",
			"Password": "123456",
		})
}

func (p *ControllerLogin) Post(ctx *gin.Context) {
	var (
		req = new(models.SUserLogin)
		c   = &Context{Context: ctx}
		err error
	)
	if err = c.ShouldBind(req); err != nil {
		c.RespErrInvalidParams(err)
		return
	}
	// check user
	recordUser := &models.SUser{Username: req.Username}
	var has bool
	if has, err = recordUser.SelectOne(recordUser); err != nil {
		c.RespErrDBError(err)
		return
	}
	if !has ||
		recordUser.Password != utils.GenPassword(req.Password, recordUser.Salt) {
		c.RespErrInvalidParams()
		return
	}
	if recordUser.State == false {
		c.RespErrForbidden()
		return
	}
	// check role
	recordRole := new(models.SRole)
	if has, err = recordRole.SelectOneByUserId(recordUser.Id); err != nil {
		c.RespErrDBError(err)
		return
	}
	if !has || recordRole.State == false {
		c.RespErrForbidden()
		return
	}
	// check success
	if err = c.SessionCreate(recordUser.Id, recordRole.Id); err != nil {
		c.RespErrInternalServerError(err)
		return
	}
	// cache user
	SetCacheUser(recordUser)
	c.Log(log.TypeLogin, log.LevelInfo)
	c.RespRedirect302("/admin")
}
