package admin

import (
	"ginserver/internal/app/models/log"
	"ginserver/internal/app/models/s_user"
	"ginserver/tools/utils"

	"github.com/gin-gonic/gin"
)

type ControllerChangePwd struct {
	Controller
}

func (p *ControllerChangePwd) Get(ctx *gin.Context) {
	p.New(ctx)
	p.Render("change_pwd.tpl", nil)
}

func (p *ControllerChangePwd) Post(ctx *gin.Context) {
	p.New(ctx)
	var (
		req s_user.UserChangePwd
		err error
	)
	if err = ctx.ShouldBind(&req); err != nil {
		p.RespErrInvalidParams()
		return
	}
	if req.Password == req.NewPassword {
		p.RespErrInvalidParams("Old password is the same as the current password!")
		return
	}
	if req.NewPassword != req.ConfirmPassword {
		p.RespErrInvalidParams("Confirm password and new password do not match.")
		return
	}
	// TODO: verify password strength
	var recordUser = s_user.UserUpdate{Id: p.GetUserId()}
	if _, err := recordUser.SelectOne(&recordUser); err != nil {
		p.RespErrDBError(err)
		return
	}
	if recordUser.Password != utils.GenPassword(req.Password, recordUser.Salt) {
		p.RespErrInvalidParams("Password incorrect.")
		return
	}
	recordUser.Password = utils.GenPassword(req.NewPassword, recordUser.Salt)
	if _, err = recordUser.Update(&recordUser, recordUser.Id, "password"); err != nil {
		p.RespErrDBError(err)
		return
	}
	p.LogDB(log.TypeChangePwd, log.LevelInfo)
	p.RespCreated(nil)
}
