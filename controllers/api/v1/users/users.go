package users

import (
	"strconv"

	"ginserver/controllers/api"
	"ginserver/models/s_user"
	"ginserver/pkg/db"
	"ginserver/tools/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ControllerUsers struct {
	api.Controller
}

// GetOne godoc
// @Summary get an user record
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "user id"
// @Success 200 {object} s_user.User
// @Failure 400 {object} e.ResponseErr
// @Failure 404 {object} e.ResponseErr
// @Failure 501 {object} e.ResponseErr
// @Router /users/{id} [get]
func (p *ControllerUsers) GetOne(ctx *gin.Context) {
	p.New(ctx)
	id, _ := strconv.Atoi(ctx.Param("id"))
	if id < 1 {
		p.RespErrInvalidParams()
		return
	}
	cols, _ := ctx.GetQueryArray("cols")
	var record = s_user.User{Id: id}
	has, err := record.SelectOne(&record, cols...)
	if err != nil {
		p.RespErrDBError(err)
		logrus.Error(err)
		return
	}
	if !has {
		p.RespErrNotFound()
		return
	}
	p.RespOk(record)
}

func (p *ControllerUsers) Get(ctx *gin.Context) {
	p.New(ctx)
	var err error
	orderBy := db.NewOrderBy(ctx)
	if err = orderBy.Parse(); err != nil {
		p.RespErrInvalidParams(err)
		return
	}
	var (
		cols    = ctx.GetStringSlice("cols")
		record  s_user.User
		records []*s_user.User
	)
	if err = record.SelectCond(&record, &records, nil, orderBy.String(), db.NewPaging(ctx), cols...); err != nil {
		p.RespErrDBError(err)
		logrus.Error(err)
		return
	}
	p.RespOk(&records)
}

func (p *ControllerUsers) Post(ctx *gin.Context) {
	p.New(ctx)
	var record s_user.UserInsert
	if err := ctx.BindJSON(record); err != nil {
		p.RespErrInvalidParams(err)
		logrus.Error(err)
		return
	}

	record.Salt = utils.GenSalt()
	record.Password = utils.Md5(record.Password + record.Salt)
	record.RegisterIp = ctx.ClientIP()

	_, err := record.InsertOne(&record)
	if err != nil {
		p.RespErrDBError(err)
		logrus.Error(err)
		return
	}
	p.RespCreated(nil)
}

func (p *ControllerUsers) Put(ctx *gin.Context) {
	p.New(ctx)
	id, _ := strconv.Atoi(ctx.Param("id"))
	if id < 1 {
		p.RespErrInvalidParams(ctx)
		return
	}
	var record = s_user.UserUpdate{Id: id}
	has, err := record.IsExists(&record)
	if err != nil {
		p.RespErrDBError(err)
		logrus.Error(err)
		return
	}
	if !has {
		p.RespErrNotFound()
		return
	}

	if err := ctx.BindJSON(&record); err != nil {
		p.RespErrInvalidParams(err)
		return
	}
	record.Salt = utils.GenSalt()
	record.Password = utils.Md5(record.Password + record.Salt)
	if _, err = record.Update(&record, record.Id); err != nil {
		p.RespErrDBError(err)
		logrus.Error(err)
		return
	}
	p.RespOk(nil)
}

func (p *ControllerUsers) Delete(ctx *gin.Context) {
	p.New(ctx)
	id, _ := strconv.Atoi(ctx.Param("id"))
	if id < 1 {
		p.RespErrInvalidParams(ctx)
		return
	}

	var record = s_user.User{Id: id}
	if _, err := record.Delete(&record); err != nil {
		p.RespErrDBError(err)
		logrus.Error(err)
		return
	}
	p.RespOk(nil)
}
