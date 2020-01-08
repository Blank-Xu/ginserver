package system

import (
	"strconv"
	"time"

	"ginserver/models/system/user"
	"ginserver/pkg/context"
	"ginserver/tools/utils"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type UserController struct{}

// GetOne godoc
// @Summary get user record
// @Description get user record by id
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
		log.Err(err).Msg("select user failed")

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
// @Description get paging user records by condition
// @Accept  application/x-www-form-urlencoded
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
	if err := ctx.Bind(&record); err != nil {
		ctx.AbortResponseInvalidParams(err)
		return
	}

	records := make([]*user.User, 0, pageSize)
	if err := record.Select(&record, &records, nil, page, pageSize, cols...); err != nil {
		log.Err(err).Msg("select user failed")
		ctx.AbortResponseDatabaseErr(err)
		return
	}

	totalCount, err := record.Count(&record)
	if err != nil {
		log.Err(err).Msg("select user count failed")
		ctx.AbortResponseDatabaseErr(err)
		return
	}

	ctx.ResponsePagingOK(page, pageSize, totalCount, records)
}

type UsersGetRequest struct {
	Username     string    `form:"username" json:"username"`
	State        int       `form:"state" json:"state"`
	Nickname     string    `form:"nickname" json:"nickname"`
	Email        string    `form:"email" json:"email"`
	Phone        string    `form:"phone" json:"phone"`
	Remark       string    `form:"remark" json:"remark"`
	CreatedStart time.Time `form:"created_start" json:"created_start"`
	CreatedEnd   time.Time `form:"created_end" json:"created_end"`
	UpdatedStart time.Time `form:"updated_start" json:"updated_start"`
	UpdatedEnd   time.Time `form:"updated_end" json:"updated_end"`
	DeletedStart time.Time `form:"deleted_start" json:"deleted_start"`
	DeletedEnd   time.Time `form:"deleted_end" json:"deleted_end"`
	RegisterIp   string    `form:"register_ip" json:"register_ip"`
	// LoginTimeStart time.Time `form:"login_time_start" json:"login_time_start"`
	// LoginTimeEnd   time.Time `form:"login_time_end" json:"login_time_end"`
	// LoginIp        string    `json:"login_ip"`
}

type UsersGetResponse struct {
	Paging  context.Paging `json:"paging"`
	Records []*user.User   `json:"records"`
}

// Post godoc
// @Summary post user
// @Description post user record
// @Accept  json
// @Produce  json
// @Param cols path string false "cols"
// @Success 201 {object} system.UsersGetResponse
// @Failure 501 {object} context.Response
// @Router /users [post]
func (p *UserController) Post(c *gin.Context) {
	ctx := context.New(c)

	var record user.Insert
	if err := ctx.BindJSON(record); err != nil {
		ctx.AbortResponseInvalidParams(err)
		return
	}

	record.Salt = utils.GenSalt()
	record.Password = utils.Md5(record.Password + record.Salt)
	record.RegisterIp = ctx.ClientIP()

	_, err := record.InsertOne(&record)
	if err != nil {
		log.Err(err).Msg("insert user failed")
		ctx.AbortResponseDatabaseErr(err)
		return
	}

	ctx.ResponseDataCreated(nil)
}

// Put godoc
// @Summary put user
// @Description put user record
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "user id"
// @Param cols path string false "cols"
// @Success 204 ""
// @Failure 400 {object} context.Response
// @Failure 404 {object} context.Response
// @Failure 501 {object} context.Response
// @Router /users [put]
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
		log.Err(err).Msg("check user exist failed")
		ctx.AbortResponseDatabaseErr(err)
		return
	}
	if !has {
		ctx.AbortResponseNotFound()
		return
	}

	if err := ctx.Bind(&record); err != nil {
		ctx.AbortResponseInvalidParams(err)
		return
	}

	record.Salt = utils.GenSalt()
	record.Password = utils.Md5(record.Password + record.Salt)
	if _, err = record.Update(&record, record.Id); err != nil {
		log.Err(err).Msg("update user failed")
		ctx.AbortResponseDatabaseErr(err)
		return
	}

	ctx.ResponseNoContent()
}

// Delete godoc
// @Summary delete user
// @Description delete user by id
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "user id"
// @Success 204 ""
// @Failure 400 {object} context.Response
// @Failure 404 {object} context.Response
// @Failure 501 {object} context.Response
// @Router /users/{id} [delete]
func (p *UserController) Delete(c *gin.Context) {
	ctx := context.New(c)

	id, _ := strconv.Atoi(ctx.Param("id"))
	if id < 1 {
		ctx.AbortResponseInvalidParams()
		return
	}

	record := user.User{Id: id}
	count, err := record.Delete(&record)
	if err != nil {
		log.Err(err).Msg("delete user failed")
		ctx.AbortResponseDatabaseErr(err)
		return
	}
	if count == 0 {
		ctx.AbortResponseNotFound()
		return
	}

	ctx.ResponseNoContent()
}
