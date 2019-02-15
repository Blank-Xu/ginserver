package db

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Paging struct {
	*gin.Context

	Page  int // param page
	Limit int // param size

	Offset int
}

func NewPaging(c *gin.Context) *Paging {
	paging := &Paging{Context: c}
	paging.Parse()
	return paging
}

func (p *Paging) Parse() {
	p.Page, _ = strconv.Atoi(p.DefaultQuery("page", "1"))
	p.Limit, _ = strconv.Atoi(p.DefaultQuery("size", "5"))
	if p.Limit > 100 {
		p.Limit = 100
	}
	p.Offset = (p.Page - 1) * p.Limit
}

func (p *Paging) LimitOffset() (int, int) {
	return p.Limit, p.Offset
}
