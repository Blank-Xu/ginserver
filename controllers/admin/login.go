package admin

import (
	"net/http"
	"time"

	"ginserver/global"
	"ginserver/models/log"
	"ginserver/models/s_role"
	"ginserver/models/s_user"
	"ginserver/pkg/db"
	"ginserver/pkg/middlewares"
	"ginserver/tools/utils"

	"github.com/gin-gonic/gin"
)

type ControllerLogin struct {
	Controller
}

func (p *ControllerLogin) Get(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html",
		map[string]string{
			"AppName": global.AppName,
			"Version": global.Version,
			"Title":   "Login",
		})
}

func (p *ControllerLogin) Post(ctx *gin.Context) {
	p.New(ctx)
	var (
		req s_user.UserLogin
		err error
	)
	if err = p.ShouldBind(&req); err != nil {
		p.RespErrInvalidParams()
		return
	}
	// check user
	var (
		recordUser = s_user.User{Username: req.Username}
		has        bool
	)
	if has, err = recordUser.SelectOne(&recordUser); err != nil {
		p.RespErrDBError(err)
		return
	}
	if !has ||
		recordUser.Password != utils.Md5(req.Password+recordUser.Salt) {
		p.RespErrInvalidParams()
		return
	}
	if recordUser.State == false {
		p.RespErrForbidden()
		return
	}
	// check role
	var recordRole s_role.Role
	if has, err = recordRole.SelectOneByUserId(recordUser.Id); err != nil {
		p.RespErrDBError(err)
		return
	}
	if !has || recordRole.State == false {
		p.RespErrForbidden()
		return
	}
	// check success
	if err = middlewares.SessionCreate(ctx, recordUser.Id, recordRole.Id); err != nil {
		p.RespErrInternalServerError(err)
		return
	}
	recordUser.LoginTime = db.JSONTime(time.Now())
	recordUser.LoginIp = p.ClientIP()
	if _, err = recordUser.Update(&recordUser, recordUser.Id, "login_time,login_ip"); err != nil {
		p.RespErrDBError(err)
		return
	}
	// cache user
	SetCacheUser(recordUser)

	p.userId = recordUser.Id
	p.roleId = recordRole.Id

	p.LogDB(log.TypeLogin, log.LevelInfo)
	p.RespOk(nil)
}
