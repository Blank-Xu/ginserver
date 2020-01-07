package system

import (
	"strconv"

	"ginserver/models/system/user"
	"ginserver/pkg/context"
	"ginserver/tools/utils"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type UserController struct{}

// GetOne godoc
// @Summary get an user record
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "user id"
// @Param cols path string false "cols"
// @Success 200 {object} user.User
// @Failure 400 {object} context.Response
// @Failure 404 {object} context.Response
// @Failure 501 {object} context.Response
// @Router /users/{id} [get]
func (p *UserController) GetOne(c *gin.Context) {
	ctx := context.New(c)

	id, _ := strconv.Atoi(ctx.Param("id"))
	if id < 1 {
		ctx.AbortResponseInvalidParams()
		return
	}

	cols, _ := ctx.GetQueryArray("cols")
	record := user.User{Id: id}
	has, err := record.SelectOne(&record, cols...)
	if err != nil {
		log.Err(err)

		ctx.AbortResponseDatabaseErr(err)
		return
	}

	if !has {
		ctx.AbortResponseNotFound()
		return
	}

	ctx.ResponseDataOK(record)
}

// Get godoc
// @Summary get user records
// @Description get user records
// @Accept  json
// @Produce  json
// @Param cols path string false "cols"
// @Param page path int false "page"
// @Param page_size path int false "page_size"
// @Success 200 {object} system.UsersGetResponse
// @Failure 501 {object} context.Response
// @Router /users [get]
func (p *UserController) Get(c *gin.Context) {
	ctx := context.New(c)

	page, pageSize := ctx.GetPage()
	cols := ctx.GetStringSlice("cols")

	var record user.User
	records := make([]*user.User, 0, pageSize)
	if err := record.Select(&record, &records, page, pageSize, cols...); err != nil {
		log.Err(err)

		ctx.AbortResponseDatabaseErr(err)
		return
	}

	totalCount, err := record.Count(&record)
	if err != nil {
		log.Err(err)

		ctx.AbortResponseDatabaseErr(err)
		return
	}

	ctx.ResponsePagingOK(page, len(records), totalCount, records)
}

type UsersGetResponse struct {
	Paging  context.Paging `json:"paging"`
	Records []*user.User   `json:"records"`
}

func (p *UserController) Post(c *gin.Context) {
	ctx := context.New(c)

	var record user.Insert
	if err := ctx.BindJSON(record); err != nil {
		log.Err(err)

		ctx.AbortResponseInvalidParams(err)
		return
	}

	record.Salt = utils.GenSalt()
	record.Password = utils.Md5(record.Password + record.Salt)
	record.RegisterIp = ctx.ClientIP()

	_, err := record.InsertOne(&record)
	if err != nil {
		log.Err(err)

		ctx.AbortResponseDatabaseErr(err)
		return
	}

	ctx.ResponseDataCreated(nil)
}

func (p *UserController) Put(c *gin.Context) {
	ctx := context.New(c)

	id, _ := strconv.Atoi(ctx.Param("id"))
	if id < 1 {
		ctx.AbortResponseInvalidParams()
		return
	}

	record := user.Update{Id: id}
	has, err := record.IsExists(&record)
	if err != nil {
		log.Err(err)
		ctx.AbortResponseDatabaseErr(err)
		return
	}
	if !has {
		ctx.AbortResponseNotFound()
		return
	}

	if err := ctx.BindJSON(&record); err != nil {
		ctx.AbortResponseInvalidParams(err)
		return
	}

	record.Salt = utils.GenSalt()
	record.Password = utils.Md5(record.Password + record.Salt)
	if _, err = record.Update(&record, record.Id); err != nil {
		log.Err(err)

		ctx.AbortResponseDatabaseErr(err)
		return
	}
	ctx.ResponseDataOK(nil)
}

func (p *UserController) Delete(c *gin.Context) {
	ctx := context.New(c)

	id, _ := strconv.Atoi(ctx.Param("id"))
	if id < 1 {
		ctx.AbortResponseInvalidParams()
		return
	}

	record := user.User{Id: id}
	if _, err := record.Delete(&record); err != nil {
		log.Err(err)
		ctx.AbortResponseDatabaseErr(err)
		return
	}

	ctx.ResponseDataOK(nil)
}
