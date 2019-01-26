package base

type Conditional struct {
	Page   int64
	Size   int64
	SortBy []string // cols
	Order  []string // ASC, DESC
}

func (p *Conditional) GetPagingOffset() (limit, offset int64) {
	if p.Size > 100 {
		p.Size = 100
	}
	if p.Page < 1 {
		p.Page = 1
	}
	return p.Size, (p.Page - 1) * p.Size
}

func (p *Conditional) GetPagingId() (id, limit int64) {
	return p.Page, p.Size
}

// func (p *Conditional) GetOrderByCond() string {
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
