package admin

import (
	"encoding/json"
	"net/http"

	"ginserver/tools/middleware"

	"ginserver/init/config"
	"ginserver/internal/app/models"
	"ginserver/tools/e"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	*gin.Context
}

func (p *Controller) New(context *gin.Context) {
	p.Context = context
}

func (p *Controller) GetUserId() int {
	return p.GetInt(middleware.KeyUserId)
}

func (p *Controller) GetRoleId() int {
	return p.GetInt(middleware.KeyRoleId)
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
	p.Error(err)
	p.AbortWithStatusJSON(e.RespErrHttp(http.StatusInternalServerError))
}

func (p *Controller) RespErrDBError(err error) {
	p.Error(err)
	if gin.Mode() != gin.ReleaseMode {
		p.AbortWithStatusJSON(http.StatusNotImplemented, e.RespErrCode(e.CodeDBErr, err))
	} else {
		p.AbortWithStatusJSON(http.StatusNotImplemented, e.RespErrCode(e.CodeDBErr))
	}
}

func (p *Controller) Render(tpl string, value map[string]interface{}) {
	user, err := GetCacheUser(p.GetUserId())
	if err != nil {
		p.RespErrInternalServerError(err)
		return
	}
	menu, err := GetCacheRoleMenu(p.GetRoleId())
	if err != nil {
		p.RespErrInternalServerError(err)
		return
	}
	if value == nil {
		value = make(map[string]interface{})
	}
	value["AppName"] = config.GetConfig().AppName
	value["main_user"] = user
	value["main_menu"] = menu
	value["active_path"] = p.Request.URL.Path
	p.HTML(http.StatusOK, tpl, value)
}

var logWithoutParamsRouter = map[string]bool{
	"/admin/login":      true,
	"/admin/change_pwd": true,
}

func (p *Controller) Log(lType models.LogType, level models.LogLevel, remark ...string) {
	recordLog := &models.Log{
		Type:   lType,
		Level:  level,
		UserId: p.GetInt(middleware.KeyUserId),
		RoleId: p.GetInt(middleware.KeyRoleId),
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
