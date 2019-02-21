package users

import (
	"strconv"

	"ginserver/internal/app/controllers/api"
	"ginserver/internal/app/models"
	"ginserver/tools/db"
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
// @Success 200 {object} models.SUser
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
	record := models.NewSUser(id)
	has, err := record.SelectOne(record, cols...)
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
	record := new(models.SUser)
	cols := ctx.GetStringSlice("cols")
	var records []*models.SUser
	if err = record.SelectCond(record, &records, nil, orderBy.String(), db.NewPaging(ctx), cols...); err != nil {
		p.RespErrDBError(err)
		logrus.Error(err)
		return
	}
	p.RespOk(&records)
}

func (p *ControllerUsers) Post(ctx *gin.Context) {
	p.New(ctx)
	record := new(models.SUserInsert)
	if err := ctx.BindJSON(record); err != nil {
		p.RespErrInvalidParams(err)
		logrus.Error(err)
		return
	}

	record.Salt = utils.GenSalt()
	record.Password = utils.GenPassword(record.Password, record.Salt)
	record.RegisterIp = ctx.ClientIP()

	_, err := record.InsertOne(record)
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
	record := &models.SUserUpdate{Id: id}
	has, err := record.IsExists(record)
	if err != nil {
		p.RespErrDBError(err)
		logrus.Error(err)
		return
	}
	if !has {
		p.RespErrNotFound()
		return
	}

	if err := ctx.BindJSON(record); err != nil {
		p.RespErrInvalidParams(err)
		return
	}
	record.Salt = utils.GenSalt()
	record.Password = utils.GenPassword(record.Password, record.Salt)
	if _, err = record.Update(record, record.Id); err != nil {
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

	record := models.NewSUser(id)
	_, err := record.Delete(record)
	if err != nil {
		p.RespErrDBError(err)
		logrus.Error(err)
		return
	}
	p.RespOk(nil)
}
