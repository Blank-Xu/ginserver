package v1

import (
	"ginserver/models/system/user"
	"ginserver/pkg/context"
	"ginserver/tools/utils"

	"github.com/gin-gonic/gin"
)

func ModifyPasswordPost(c *gin.Context) {
	ctx := context.New(c)

	var (
		req user.ChangePwd
		err error
	)
	if err = ctx.ShouldBind(&req); err != nil {
		ctx.AbortResponseInvalidParams(err)
		return
	}
	if req.Password == req.NewPassword {
		ctx.AbortResponseInvalidParams("Old password is the same as the current password!")
		return
	}
	if req.NewPassword != req.ConfirmPassword {
		ctx.AbortResponseInvalidParams("Confirm password and new password do not match.")
		return
	}
	// TODO: verify password strength
	recordUser := user.Update{Id: 0} // p.GetUserId()}
	if _, err := recordUser.SelectOne(&recordUser); err != nil {
		ctx.AbortResponseDatabaseErr(err)
		return
	}
	if recordUser.Password != utils.Md5(req.Password+recordUser.Salt) {
		ctx.AbortResponseInvalidParams("Password incorrect.")
		return
	}
	recordUser.Password = utils.Md5(req.NewPassword + recordUser.Salt)
	if _, err = recordUser.Update(&recordUser, recordUser.Id, "password"); err != nil {
		ctx.AbortResponseDatabaseErr(err)
		return
	}
	// p.LogDB(log.TypeChangePwd, zerolog.InfoLevel)

	ctx.ResponseNoContent()
}
