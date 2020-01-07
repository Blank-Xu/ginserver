package context

import (
	"net/http"
)

func (p *Context) ResponseDataOK(data interface{}) {
	p.JSON(
		http.StatusOK,
		data,
	)
}

func (p *Context) AbortResponseDataOK(data interface{}) {
	p.AbortWithStatusJSON(
		http.StatusOK,
		data,
	)
}

func (p *Context) ResponsePagingOK(page, pageSize int, totalCount int64, records interface{}) {
	p.JSON(
		http.StatusOK,
		ResponsePaging{
			Paging: Paging{
				Page:       page,
				PageSize:   pageSize,
				TotalCount: totalCount,
			},
			Records: records,
		},
	)
}

func (p *Context) AbortResponsePagingOK(page, pageSize int, totalCount int64, records interface{}) {
	p.AbortWithStatusJSON(
		http.StatusOK,
		ResponsePaging{
			Paging: Paging{
				Page:       page,
				PageSize:   pageSize,
				TotalCount: totalCount,
			},
			Records: records,
		},
	)
}

func (p *Context) ResponseDataCreated(data interface{}) {
	p.JSON(
		http.StatusCreated,
		data,
	)
}

func (p *Context) AbortResponseDataCreated(data interface{}) {
	p.AbortWithStatusJSON(
		http.StatusCreated,
		data,
	)
}

func (p *Context) ResponseDataAccepted(data interface{}) {
	p.JSON(
		http.StatusAccepted,
		data,
	)
}

func (p *Context) AbortResponseDataAccepted(data interface{}) {
	p.AbortWithStatusJSON(
		http.StatusAccepted,
		data,
	)
}

func (p *Context) ResponseNoContent() {
	p.Status(http.StatusNoContent)
}

func (p *Context) AbortResponseNoContent() {
	p.AbortWithStatus(http.StatusNoContent)
}
