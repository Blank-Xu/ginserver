package users

import (
	"strconv"

	"ginserver/internal/app/models"
	"ginserver/tools/db"
	"ginserver/tools/e"
	"ginserver/tools/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ControllerUsers struct{}

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
// @Router /admins/{id} [get]
func (p *ControllerUsers) GetOne(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if id < 1 {
		e.RespErrInvalidParams(ctx)
		return
	}
	cols, _ := ctx.GetQueryArray("cols")
	record := models.NewSUser(id)
	has, err := record.SelectOne(record, cols...)
	if err != nil {
		e.RespErrDBError(ctx, err)
		logrus.Error(err)
		return
	}
	if !has {
		e.RespErrNotFound(ctx)
		return
	}
	e.RespDataOk(ctx, record)
}

func (p *ControllerUsers) Get(ctx *gin.Context) {
	var err error
	orderBy := db.NewOrderBy(ctx)
	if err = orderBy.Parse(); err != nil {
		e.RespErrInvalidParams(ctx, err)
		return
	}
	record := new(models.SUser)
	cols := ctx.GetStringSlice("cols")
	var records []*models.SUser
	if err = record.SelectCond(record, &records, nil, orderBy.String(), db.NewPaging(ctx), cols...); err != nil {
		e.RespErrDBError(ctx, err)
		logrus.Error(err)
		return
	}
	e.RespDataOk(ctx, &records)
}

func (p *ControllerUsers) Post(ctx *gin.Context) {
	record := new(models.SUserInsert)
	if err := ctx.BindJSON(record); err != nil {
		e.RespErrInvalidParams(ctx, err)
		logrus.Error(err)
		return
	}

	record.Salt = utils.GenSalt()
	record.Password = utils.GenPassword(record.Password, record.Salt)
	record.RegisterIp = ctx.ClientIP()

	count, err := record.InsertOne(record)
	if err != nil {
		e.RespErrDBError(ctx, err)
		logrus.Error(err)
		return
	}
	e.RespDataCreated(ctx, count)
}

func (p *ControllerUsers) Put(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if id < 1 {
		e.RespErrInvalidParams(ctx)
		return
	}
	record := &models.SUserUpdate{Id: id}
	has, err := record.IsExists(record)
	if err != nil {
		e.RespErrDBError(ctx, err)
		logrus.Error(err)
		return
	}
	if !has {
		e.RespErrNotFound(ctx)
		return
	}

	if err := ctx.BindJSON(record); err != nil {
		e.RespErrInvalidParams(ctx, err)
		return
	}
	record.Salt = utils.GenSalt()
	record.Password = utils.GenPassword(record.Password, record.Salt)
	count, err := record.Update(record, record.Id)
	if err != nil {
		e.RespErrDBError(ctx, err)
		logrus.Error(err)
		return
	}
	e.RespDataOk(ctx, count)
}

func (p *ControllerUsers) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if id < 1 {
		e.RespErrInvalidParams(ctx)
		return
	}

	record := models.NewSUser(id)
	count, err := record.Delete(record)
	if err != nil {
		e.RespErrDBError(ctx, err)
		logrus.Error(err)
		return
	}
	e.RespDataOk(ctx, count)
}
