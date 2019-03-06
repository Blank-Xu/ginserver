package admin

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"

	"ginserver/internal/app/models/log"

	"ginserver/tools/middleware"

	"ginserver/init/config"
	"ginserver/tools/e"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	*gin.Context
	userId int
	roleId int
}

// New need set context first
func (p *Controller) New(context *gin.Context) {
	p.Context = context
	p.userId = p.GetInt(middleware.KeyUserId)
	p.roleId = p.GetInt(middleware.KeyRoleId)
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
	p.AbortWithStatusJSON(http.StatusBadRequest, e.RespErrCode(e.CodeInvalidParams, err...))
}

func (p *Controller) RespErrForbidden() {
	p.AbortWithStatusJSON(e.RespErrHttp(http.StatusForbidden))
}

func (p *Controller) RespErrNotFound() {
	p.AbortWithStatusJSON(e.RespErrHttp(http.StatusNotFound))
}

func (p *Controller) RespErrInternalServerError(err error) {
	p.LogDB(log.TypeDBError, log.LevelError, p.Error(err).Error())

	p.AbortWithStatusJSON(e.RespErrHttp(http.StatusInternalServerError))
}

func (p *Controller) RespErrDBError(err error) {
	p.LogDB(log.TypeInternalServerError, log.LevelError, p.Error(err).Error())

	if gin.Mode() != gin.ReleaseMode {
		p.AbortWithStatusJSON(http.StatusNotImplemented, e.RespErrCode(e.CodeDBErr, err))
	} else {
		p.AbortWithStatusJSON(http.StatusNotImplemented, e.RespErrCode(e.CodeDBErr))
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
	value["AppName"] = config.GetConfig().AppName
	value["MainUser"] = user
	value["MainMenu"] = menu
	value["ActivePath"] = p.Request.URL.Path
	p.HTML(http.StatusOK, tpl, value)
}

func (p *Controller) LogErr(err error) {
	if err != nil {
		logrus.Error(err)
	}
}

var logWithoutParamsPath = []string{
	"/admin/login",
	"/admin/change_pwd",
}

func (p *Controller) LogDB(lType log.Type, level log.Level, remark ...string) {
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
