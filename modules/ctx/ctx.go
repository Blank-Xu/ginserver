package ctx

import "github.com/gin-gonic/gin"

type Ctx struct {
	*gin.Context

	Uid  int
	Role string

	Cols   []string
	Page   int64
	Size   int64
	SortBy []string // cols
	Order  []string // ASC, DESC
}

func (p *Ctx) Parse() error {
	return nil
}

func (p *Ctx) GetPagingOffset() (limit, offset int64) {
	if p.Size > 100 {
		p.Size = 100
	}
	if p.Page < 1 {
		p.Page = 1
	}
	return p.Size, (p.Page - 1) * p.Size
}

func (p *Ctx) GetPagingId() (id, limit int64) {
	return p.Page, p.Size
}

// func (p *Ctx) GetOrderByCond() string {
// 	lenSortBy := len(p.SortBy)
// 	if lenSortBy == 0 {
// 		return ""
// 	}
// 	lenOrder := len(p.Order)
// 	if lenOrder == 0 {
// 		p.Order[0] = " DESC "
// 	}
// 	return ""
// }

func (p *Ctx) IsLogin() bool {
	return (p.Uid > 0) && (len(p.Role) > 0)
}
