package db

import (
	"github.com/gin-gonic/gin"
)

type Paging struct {
	*gin.Context

	Page  int
	Limit int

	Id     int
	Offset int
}

func NewPaging(c *gin.Context) *Paging {
	paging := &Paging{Context: c}
	paging.Parse()
	return paging
}

func (p *Paging) Parse() {
	p.Page = p.GetInt("page")
	p.Limit = p.GetInt("size")
	if p.Limit > 100 {
		p.Limit = 100
	} else if p.Limit < 1 {
		p.Limit = 5
	}
	if p.Page < 1 {
		p.Page = 1
	}
	p.Offset = (p.Page - 1) * p.Limit
}

func (p *Paging) LimitOffset() (int, int) {
	return p.Limit, p.Offset
}

func (p *Paging) IdLimit() (int, int) {
	return p.Page, p.Limit
}
