package user_set

import (
	"ginserver/models"
	"ginserver/modules/utils"

	"github.com/gin-gonic/gin"

	"ginserver/controllers/admin"
)

type ControllerChangePwd struct{}

func (p *ControllerChangePwd) Get(ctx *gin.Context) {
	newCtx := admin.NewContext(ctx)
	newCtx.Render("user_set/change_pwd.tpl", gin.H{})
}

func (p *ControllerChangePwd) Post(ctx *gin.Context) {
	var err error
	newCtx := admin.NewContext(ctx)
	req := new(reqChangePwd)
	if err = ctx.ShouldBind(req); err != nil {
		newCtx.RespErrInvalidParams()
		return
	}
	if req.Password == req.NewPassword {
		newCtx.RespErrInvalidParams("Old password is the same as the current password!") // 旧密码与当前密码相同！
		return
	}
	if req.NewPassword != req.ConfirmPassword {
		newCtx.RespErrInvalidParams("Confirm password and new password do not match.") // 您的确认密码和新密码不一致。
		return
	}
	// TODO: 验证密码强度
	recordUser := &models.SUserUpdate{Id: newCtx.GetUserId()}
	if _, err := recordUser.SelectOne(recordUser); err != nil {
		newCtx.RespErrDBError(err)
		return
	}
	if recordUser.Password != utils.GenPassword(req.Password, recordUser.Salt) {
		newCtx.RespErrInvalidParams("Password incorrect.")
		return
	}
	recordUser.Password = utils.GenPassword(req.NewPassword, recordUser.Salt)
	if _, err = recordUser.Update(recordUser); err != nil {
		newCtx.RespErrDBError(err)
		return
	}
	newCtx.RespDataAccepted(nil)
}

type reqChangePwd struct {
	Password        string `form:"password" binding:"required"`
	NewPassword     string `form:"new_password" binding:"required"`
	ConfirmPassword string `form:"confirm_password" binding:"required"`
}
