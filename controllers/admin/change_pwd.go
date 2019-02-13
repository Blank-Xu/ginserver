package admin

import (
	"ginserver/models"
	"ginserver/modules/utils"

	"github.com/gin-gonic/gin"
)

type ControllerChangePwd struct{}

func (p *ControllerChangePwd) Get(ctx *gin.Context) {
	c := ContextParse(ctx)
	if c == nil {
		return
	}
	c.Render("change_pwd.tpl", gin.H{})
}

func (p *ControllerChangePwd) Post(ctx *gin.Context) {
	c := ContextParse(ctx)
	if c == nil {
		return
	}
	var err error
	req := new(reqChangePwd)
	if err = ctx.ShouldBind(req); err != nil {
		c.RespErrInvalidParams()
		return
	}
	if req.Password == req.NewPassword {
		c.RespErrInvalidParams("Old password is the same as the current password!") // 旧密码与当前密码相同！
		return
	}
	if req.NewPassword != req.ConfirmPassword {
		c.RespErrInvalidParams("Confirm password and new password do not match.") // 您的确认密码和新密码不一致。
		return
	}
	// TODO: 验证密码强度
	recordUser := &models.SUserUpdate{Id: c.GetUserId()}
	if _, err := recordUser.SelectOne(recordUser); err != nil {
		c.RespErrDBError(err)
		return
	}
	if recordUser.Password != utils.GenPassword(req.Password, recordUser.Salt) {
		c.RespErrInvalidParams("Password incorrect.")
		return
	}
	recordUser.Password = utils.GenPassword(req.NewPassword, recordUser.Salt)
	if _, err = recordUser.Update(recordUser); err != nil {
		c.RespErrDBError(err)
		return
	}
	c.RespDataAccepted(nil)
}

type reqChangePwd struct {
	Password        string `form:"password" binding:"required"`
	NewPassword     string `form:"new_password" binding:"required"`
	ConfirmPassword string `form:"confirm_password" binding:"required"`
}
