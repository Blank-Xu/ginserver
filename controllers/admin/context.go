package admin

import (
	"encoding/json"
	"errors"
	"net/http"

	"ginserver/modules/config"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"ginserver/models"
	"ginserver/modules/casbin"
	"ginserver/modules/e"
	"ginserver/modules/log"
)

type Context struct {
	*gin.Context
	userId int
	roleId int
}

// ParseContext to check login and casbin rules
// need to run in controller first
func (p *Context) ParseContext(context *gin.Context, enforce ...bool) bool {
	session := sessions.Default(context)
	if session != nil {
		// login session check
		vUid := session.Get("userId")
		vRole := session.Get("roleId")
		if userId, ok := vUid.(int); ok {
			if roleId, ok := vRole.(int); ok {
				if userId > 0 && roleId > 0 {
					if len(enforce) > 0 && enforce[0] == false {
						p.new(context, userId, roleId)
						return true
						// casbin rules check
					} else if casbin.Enforce(context, userId) {
						p.new(context, userId, roleId)
						return true
					}
					context.AbortWithStatusJSON(e.RespErrHttp(http.StatusForbidden))
					return false
				}
			}
		}
	}
	context.Redirect(http.StatusFound, "/admin/login")
	return false
}

func (p *Context) new(context *gin.Context, userId, roleId int) {
	p.Context = context
	p.userId = userId
	p.roleId = roleId
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

func (p *Context) Render(tpl string, value map[string]interface{}) {
	if p.userId == 0 {
		p.RespErrInvalidParams()
		return
	}
	user, err := GetCacheUser(p.userId)
	if err != nil {
		p.RespErrInternalServerError(err)
		return
	}
	menu, err := GetCacheRoleMenu(p.roleId)
	if err != nil {
		p.RespErrInternalServerError(err)
		return
	}
	value["AppName"] = config.GetConfig().AppName
	value["main_user"] = user
	value["main_menu"] = menu
	value["active_path"] = p.Request.URL.Path
	p.HTML(http.StatusOK, tpl, value)
}

var logWithoutParamsRouter = map[string]bool{
	"/admin/login": true,
}

func (p *Context) Log(lType log.Type, level log.Level, remark ...string) {
	recordLog := &models.Log{
		Type:   lType,
		Level:  level,
		UserId: p.userId,
		RoleId: p.roleId,
		Method: p.Request.Method,
		Path:   p.Request.URL.Path,
		Ip:     p.ClientIP(),
	}
	if !logWithoutParamsRouter[recordLog.Path] {
		param, _ := json.Marshal(p.Request.Form)
		recordLog.Params = string(param)
	}
	if len(remark) > 0 {
		recordLog.Remark = remark[0]
	}
	recordLog.InsertOne(recordLog)
}
