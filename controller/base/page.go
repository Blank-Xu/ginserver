package base

type Page struct {
	Page int64
	Size int64
}

func (p *Page) GetPagingOffset() (limit, offset int64) {
	if p.Size > 100 {
		p.Size = 100
	}
	if p.Page < 1 {
		p.Page = 1
	}
	return p.Size, (p.Page - 1) * p.Size
}

func (p *Page) GetPagingId() (id, limit int64) {
	return p.Page, p.Size
}
