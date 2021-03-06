package admin

import (
	"ginserver/internal/app/models/log"
	"ginserver/internal/app/models/s_user"

	"github.com/gin-gonic/gin"
)

type ControllerInfo struct {
	Controller
}

func (p *ControllerInfo) Get(ctx *gin.Context) {
	p.New(ctx)
	p.Render("info.tpl", nil)
}

func (p *ControllerInfo) Post(ctx *gin.Context) {
	p.New(ctx)
	var (
		req = s_user.UserInfoUpdate{Id: p.userId}
		err error
	)
	if err = p.ShouldBind(&req); err != nil || len(req.Nickname) == 0 {
		p.RespErrInvalidParams()
		return
	}
	req.Updater = p.userId
	if _, err = req.Update(&req, p.userId); err != nil {
		p.RespErrDBError(err)
		return
	}
	if _, err = SetCacheUserById(p.userId); err != nil {
		p.RespErrInternalServerError(err)
		return
	}
	p.LogDB(log.TypeOther, log.LevelInfo)
	p.RespOk(nil)
}
