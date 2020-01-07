package context

import (
	"net/http"
)

func (p *Context) Redirect301(location string) {
	p.Redirect(
		http.StatusMovedPermanently,
		location,
	)
}

func (p *Context) Redirect302(location string) {
	p.Redirect(
		http.StatusFound,
		location,
	)
}
