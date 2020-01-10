package v1

import (
	"time"

	"ginserver/models/system/role"
	"ginserver/models/system/user"
	"ginserver/pkg/context"
	"ginserver/pkg/middlewares"
	"ginserver/tools/utils"

	"github.com/gin-gonic/gin"
)

func LoginPost(c *gin.Context) {
	ctx := context.New(c)

	var (
		req user.Login
		err error
	)
	if err = ctx.ShouldBind(&req); err != nil {
		ctx.AbortResponseInvalidParams(err)
		return
	}

	// check user
	var (
		recordUser = user.User{Username: req.Username}
		has        bool
	)
	if has, err = recordUser.SelectOne(&recordUser); err != nil {
		ctx.AbortResponseDatabaseErr(err)
		return
	}
	if !has ||
		recordUser.Password != utils.Md5(req.Password+recordUser.Salt) {
		ctx.AbortResponseInvalidParams()
		return
	}
	if recordUser.State == false {
		ctx.AbortResponseForbidden()
		return
	}
	// check role
	var recordRole role.Role
	if has, err = recordRole.SelectOneByUserId(recordUser.Id); err != nil {
		ctx.AbortResponseDatabaseErr(err)
		return
	}
	if !has || recordRole.State == false {
		ctx.AbortResponseForbidden()
		return
	}
	// check success
	if err = middlewares.NewSession(c, recordUser.Id, recordRole.Id); err != nil {
		ctx.AbortResponseInternalServerError(err)
		return
	}
	recordUser.LoginTime = time.Now()
	recordUser.LoginIp = c.ClientIP()
	if _, err = recordUser.Update(&recordUser, recordUser.Id, "login_time,login_ip"); err != nil {
		ctx.AbortResponseDatabaseErr(err)
		return
	}

	// cache user
	SetCacheUser(recordUser)

	// p.userId = recordUser.Id
	// p.roleId = recordRole.Id
	//
	// p.LogDB(log.TypeLogin, zerolog.InfoLevel)

	ctx.ResponseNoContent()
}
