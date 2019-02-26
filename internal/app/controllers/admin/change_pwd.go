package admin

import (
	"ginserver/internal/app/models"
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
	var err error
	req := new(models.SUserChangePwd)
	if err = ctx.ShouldBind(req); err != nil {
		p.RespErrInvalidParams()
		return
	}
	if req.Password == req.NewPassword {
		p.RespErrInvalidParams("Old password is the same as the current password!") // 旧密码与当前密码相同！
		return
	}
	if req.NewPassword != req.ConfirmPassword {
		p.RespErrInvalidParams("Confirm password and new password do not match.") // 您的确认密码和新密码不一致。
		return
	}
	// TODO: 验证密码强度
	recordUser := &models.SUserUpdate{Id: p.GetUserId()}
	if _, err := recordUser.SelectOne(recordUser); err != nil {
		p.RespErrDBError(err)
		return
	}
	if recordUser.Password != utils.GenPassword(req.Password, recordUser.Salt) {
		p.RespErrInvalidParams("Password incorrect.")
		return
	}
	recordUser.Password = utils.GenPassword(req.NewPassword, recordUser.Salt)
	if _, err = recordUser.Update(recordUser, recordUser.Id, "password"); err != nil {
		p.RespErrDBError(err)
		return
	}
	p.Log(models.LogTypeChangePwd, models.LogLevelInfo)
	p.RespCreated(nil)
}
