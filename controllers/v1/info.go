package v1

import (
	"ginserver/models/system/user"
	"ginserver/pkg/context"

	"github.com/gin-gonic/gin"
)

type InfoController struct {
}

func (InfoController) Get(c *gin.Context) {
	// ctx := context.New(c)

}

func (InfoController) Post(c *gin.Context) {
	ctx := context.New(c)

	userId := 0
	var (
		req = user.InfoUpdate{Id: userId}
		err error
	)
	if err = ctx.ShouldBind(&req); err != nil || len(req.Nickname) == 0 {
		ctx.AbortResponseInvalidParams()
		return
	}
	// req.Updater = p.userId
	if _, err = req.Update(&req, userId); err != nil {
		ctx.AbortResponseDatabaseErr(err)
		return
	}
	if _, err = SetCacheUserById(userId); err != nil {
		ctx.AbortResponseInternalServerError(err)
		return
	}
	// p.LogDB(log.TypeOther, zerolog.InfoLevel)

	ctx.ResponseNoContent()
}
