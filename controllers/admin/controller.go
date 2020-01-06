package admin

import (
	"encoding/json"
	"net/http"

	"ginserver/global"
	"ginserver/models/log"
	"ginserver/pkg/middlewares"
	"ginserver/pkg/resp"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

type Controller struct {
	*gin.Context
	userId int
	roleId int
}

// New need set context first
func (p *Controller) New(context *gin.Context) {
	p.Context = context
	p.userId = p.GetInt(middlewares.KeyUserId)
	p.roleId = p.GetInt(middlewares.KeyRoleId)
}

func (p *Controller) GetUserId() int {
	return p.userId
}

func (p *Controller) GetRoleId() int {
	return p.roleId
}

// RespOk 200
// for GET
func (p *Controller) RespOk(data interface{}) {
	if data == nil {
		p.AbortWithStatus(http.StatusOK)
		return
	}
	p.AbortWithStatusJSON(http.StatusOK, data)
}

// RespCreated 201
// for POST/PUT/PATCH
func (p *Controller) RespCreated(data interface{}) {
	if data == nil {
		p.AbortWithStatus(http.StatusCreated)
		return
	}
	p.AbortWithStatusJSON(http.StatusCreated, data)
}

// RespAccepted  202
// for async task
func (p *Controller) RespAccepted() {
	p.AbortWithStatus(http.StatusAccepted)
}

// RespNoContent 204
// for DELETE
func (p *Controller) RespNoContent() {
	p.AbortWithStatus(http.StatusNoContent)
}

func (p *Controller) RespRedirect301(location string) {
	p.Redirect(http.StatusMovedPermanently, location)
}

func (p *Controller) RespRedirect302(location string) {
	p.Redirect(http.StatusFound, location)
}

func (p *Controller) RespErrInvalidParams(err ...interface{}) {
	p.AbortWithStatusJSON(http.StatusBadRequest, resp.RespErrCode(resp.CodeInvalidParams, err...))
}

func (p *Controller) RespErrForbidden() {
	p.AbortWithStatusJSON(resp.RespErrHttp(http.StatusForbidden))
}

func (p *Controller) RespErrNotFound() {
	p.AbortWithStatusJSON(resp.RespErrHttp(http.StatusNotFound))
}

func (p *Controller) RespErrInternalServerError(err error) {
	p.LogDB(log.TypeDBError, zerolog.ErrorLevel, p.Error(err).Error())

	p.AbortWithStatusJSON(resp.RespErrHttp(http.StatusInternalServerError))
}

func (p *Controller) RespErrDBError(err error) {
	p.LogDB(log.TypeInternalServerError, zerolog.ErrorLevel, p.Error(err).Error())

	if gin.Mode() != gin.ReleaseMode {
		p.AbortWithStatusJSON(http.StatusNotImplemented, resp.RespErrCode(resp.CodeDBErr, err))
	} else {
		p.AbortWithStatusJSON(http.StatusNotImplemented, resp.RespErrCode(resp.CodeDBErr))
	}
}

func (p *Controller) Render(tpl string, value map[string]interface{}) {
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
	if value == nil {
		value = make(map[string]interface{})
	}
	value["AppName"] = global.AppName
	value["MainUser"] = user
	value["MainMenu"] = menu
	value["ActivePath"] = p.Request.URL.Path
	p.HTML(http.StatusOK, tpl, value)
}

func (p *Controller) LogErr(err error) {
	if err != nil {
		zlog.Err(err)
	}
}

var logWithoutParamsPath = []string{
	"/admin/login",
	"/admin/change_pwd",
}

func (p *Controller) LogDB(lType log.Type, level zerolog.Level, remark ...string) {
	var (
		params  string
		lRemark string
	)
	for _, v := range logWithoutParamsPath {
		if v == p.Request.URL.Path {
			params = "{}"
			break
		}
	}
	if len(params) == 0 {
		if p.Request.Form != nil {
			param, _ := json.Marshal(p.Request.Form)
			params = string(param)
		} else {
			params = "{}"
		}
	}

	if len(remark) > 0 {
		lRemark = remark[0]
	}

	p.LogErr(
		log.Insert(level, lType, p.userId, p.roleId, p.Request.Method, p.Request.URL.Path,
			params, p.ClientIP(), lRemark))
}
