package admin

import (
	"errors"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/json"

	"ginserver/models"
	"ginserver/modules/e"
)

type Context struct {
	*gin.Context
	userId int
	roleId int
}

func NewContext(context *gin.Context) *Context {
	return &Context{Context: context}
}

func (p *Context) SessionCreate(userId, roleId int) error {
	session := sessions.Default(p.Context)
	if session != nil {
		session.Set("userId", userId)
		session.Set("roleId", roleId)
		p.userId = userId
		p.roleId = roleId
		return session.Save()
	}
	return errors.New("session is nil")
}

func (p *Context) SessionParse() (ok bool) {
	session := sessions.Default(p.Context)
	if session != nil {
		vUid := session.Get("userId")
		vRole := session.Get("roleId")
		if vUid != nil && vRole != nil {
			if p.userId, ok = vUid.(int); ok {
				if p.roleId, ok = vRole.(int); ok {
					return
				}
			}
		}
	}
	return
}

func (p *Context) SessionDestroy() {
	session := sessions.Default(p.Context)
	if session != nil {
		session.Clear()
		if err := session.Save(); err != nil {
			p.Error(err)
		}
	}
}

func (p *Context) GetRoleId() int {
	return p.roleId
}

func (p *Context) GetUserId() int {
	return p.userId
}

var logWithoutParamsRouter = map[string]bool{
	"/admin/login": true,
}

func (p *Context) Log(logType models.LogType, remark ...string) {
	log := &models.SLog{
		LogType: logType,
		UserId:  p.userId,
		RoleId:  p.roleId,
		Method:  p.Request.Method,
		Path:    p.Request.URL.Path,
		Ip:      p.ClientIP(),
	}
	if !logWithoutParamsRouter[log.Path] {
		param, _ := json.Marshal(p.Request.Form)
		log.Params = string(param)
	}
	if len(remark) > 0 {
		log.Remark = remark[0]
	}
	log.InsertOne(log)
}

func (p *Context) RespDataOk(data interface{}) {
	p.JSON(http.StatusOK, data)
}

func (p *Context) RespDataCreated(data interface{}) {
	p.JSON(http.StatusCreated, data)
}

func (p *Context) RespDataAccepted(data interface{}) {
	p.JSON(http.StatusAccepted, data)
}

func (p *Context) RespRedirect301(location string) {
	p.Redirect(http.StatusMovedPermanently, location)
}

func (p *Context) RespRedirect302(location string) {
	p.Redirect(http.StatusFound, location)
}

func (p *Context) RespErrInvalidParams(err ...interface{}) {
	p.AbortWithStatusJSON(http.StatusBadRequest, e.RespErrCode(e.CodeInvalidParams, err...))
}

func (p *Context) RespErrForbidden() {
	p.AbortWithStatusJSON(e.RespErrHttp(http.StatusForbidden))
}

func (p *Context) RespErrNotFound() {
	p.AbortWithStatusJSON(e.RespErrHttp(http.StatusNotFound))
}

func (p *Context) RespErrInternalServerError(err error) {
	p.Error(err)
	p.AbortWithStatusJSON(e.RespErrHttp(http.StatusInternalServerError))
}

func (p *Context) RespErrDBError(err error) {
	p.Error(err)
	if gin.Mode() != gin.ReleaseMode {
		p.AbortWithStatusJSON(http.StatusNotImplemented, e.RespErrCode(e.CodeDBErr, err))
	} else {
		p.AbortWithStatusJSON(http.StatusNotImplemented, e.RespErrCode(e.CodeDBErr))
	}
}
