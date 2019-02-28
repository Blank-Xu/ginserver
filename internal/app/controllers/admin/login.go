package admin

import (
	"net/http"
	"time"

	"ginserver/init/config"
	"ginserver/internal/app/models/log"
	"ginserver/internal/app/models/s_role"
	"ginserver/internal/app/models/s_user"
	"ginserver/tools/db"
	"ginserver/tools/middleware"
	"ginserver/tools/utils"

	"github.com/gin-gonic/gin"
)

type ControllerLogin struct {
	Controller
}

func (p *ControllerLogin) Get(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html",
		map[string]string{
			"AppName": config.GetConfig().AppName,
			"Version": config.Version,
			"Title":   "Login",
		})
}

func (p *ControllerLogin) Post(ctx *gin.Context) {
	p.New(ctx)
	var (
		req = new(s_user.UserLogin)
		err error
	)
	if err = p.ShouldBind(req); err != nil {
		p.RespErrInvalidParams()
		return
	}
	// check user
	recordUser := &s_user.User{Username: req.Username}
	var has bool
	if has, err = recordUser.SelectOne(recordUser); err != nil {
		p.RespErrDBError(err)
		return
	}
	if !has ||
		recordUser.Password != utils.GenPassword(req.Password, recordUser.Salt) {
		p.RespErrInvalidParams()
		return
	}
	if recordUser.State == false {
		p.RespErrForbidden()
		return
	}
	// check role
	recordRole := new(s_role.Role)
	if has, err = recordRole.SelectOneByUserId(recordUser.Id); err != nil {
		p.RespErrDBError(err)
		return
	}
	if !has || recordRole.State == false {
		p.RespErrForbidden()
		return
	}
	// check success
	if err = middleware.SessionCreate(ctx, recordUser.Id, recordRole.Id); err != nil {
		p.RespErrInternalServerError(err)
		return
	}
	recordUser.LoginTime = db.JSONTime(time.Now())
	recordUser.LoginIp = p.ClientIP()
	recordUser.Update(recordUser, recordUser.Id, "login_time,login_ip")
	// cache user
	SetCacheUser(recordUser)
	p.LogDB(log.TypeLogin, log.LevelInfo)
	p.RespOk(nil)
}
