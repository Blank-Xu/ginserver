package admin

import (
	"net/http"
	"time"

	"ginserver/init/config"
	"ginserver/internal/app/models"
	"ginserver/tools/db"
	"ginserver/tools/utils"

	"github.com/gin-gonic/gin"
)

type ControllerLogin struct{}

func (p *ControllerLogin) Get(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html",
		map[string]string{
			"AppName": config.GetConfig().AppName,
			"Version": config.Version,
			"Title":   "Login",
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
	if err = c.sessionCreate(recordUser.Id, recordRole.Id); err != nil {
		c.RespErrInternalServerError(err)
		return
	}
	recordUser.LoginTime = db.JSONTime(time.Now())
	recordUser.LoginIp = c.ClientIP()
	recordUser.Update(recordUser, recordUser.Id, "login_time,login_ip")
	// cache user
	SetCacheUser(recordUser)
	c.Log(models.LogTypeLogin, models.LogLevelInfo)
	c.RespRedirect302("/admin")
}
