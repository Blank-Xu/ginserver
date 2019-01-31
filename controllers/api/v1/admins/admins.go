package admins

import (
	"strconv"

	"ginserver/models"
	"ginserver/modules/db"
	"ginserver/modules/e"
	"ginserver/modules/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Admins struct{}

func (p *Admins) RegisterRouter(r *gin.RouterGroup) {
	r.GET("admins/:id", p.GetOne)
	r.GET("admins", p.Get)
	r.POST("admins", p.Post)
	r.PUT("admins/:id", p.Put)
	r.DELETE("admins/:id", p.Delete)
}

func (p *Admins) GetOne(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if id < 1 {
		e.RespErrParamsInvalid(c)
		return
	}
	cols, _ := c.GetQueryArray("cols")
	record := models.NewSAdmin(id)
	has, err := record.SelectOne(record, cols...)
	if err != nil {
		e.RespErrDBError(c, err)
		logrus.Error(err)
		return
	}
	if !has {
		e.RespErrNotFound(c)
		return
	}
	e.RespSuccOK(c, record)
}

func (p *Admins) Get(c *gin.Context) {
	var err error
	orderBy := db.NewOrderBy(c)
	if err = orderBy.Parse(); err != nil {
		e.RespErrParamsInvalid(c, err)
		return
	}
	record := new(models.SAdmin)
	cols := c.GetStringSlice("cols")
	var records []*models.SAdmin
	if err = record.SelectCond(record, &records, nil, orderBy.String(), db.NewPaging(c), cols...); err != nil {
		e.RespErrDBError(c, err)
		logrus.Error(err)
		return
	}
	e.RespSuccOK(c, &records)
}

func (p *Admins) Post(c *gin.Context) {
	record := new(models.SAdminInsert)
	if err := c.BindJSON(record); err != nil {
		e.RespErrParamsInvalid(c, err)
		logrus.Error(err)
		return
	}

	record.Salt = utils.GenSalt()
	record.Password = utils.GenPassword(record.Password, record.Salt)
	record.RegisterIp = c.ClientIP()

	count, err := record.InsertOne(record)
	if err != nil {
		e.RespErrDBError(c, err)
		logrus.Error(err)
		return
	}
	e.RespSuccCreated(c, count)
}

func (p *Admins) Put(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if id < 1 {
		e.RespErrParamsInvalid(c)
		return
	}
	record := &models.SAdminUpdate{Id: id}
	has, err := record.IsExists(record)
	if err != nil {
		e.RespErrDBError(c, err)
		logrus.Error(err)
		return
	}
	if !has {
		e.RespErrNotFound(c)
		return
	}

	if err := c.BindJSON(record); err != nil {
		e.RespErrParamsInvalid(c, err)
		return
	}
	record.Salt = utils.GenSalt()
	record.Password = utils.GenPassword(record.Password, record.Salt)
	count, err := record.Update(record)
	if err != nil {
		e.RespErrDBError(c, err)
		logrus.Error(err)
		return
	}
	e.RespSuccOK(c, count)
}

func (p *Admins) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if id < 1 {
		e.RespErrParamsInvalid(c)
		return
	}

	record := models.NewSAdmin(id)
	count, err := record.Delete(record)
	if err != nil {
		e.RespErrDBError(c, err)
		logrus.Error(err)
		return
	}
	e.RespSuccOK(c, count)
}
