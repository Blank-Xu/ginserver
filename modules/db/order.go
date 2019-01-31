package db

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	separator    = ","
	defaultOrder = " DESC"
)

type OrderBy struct {
	*gin.Context

	Fields string // order columns
	Order  string // ASC, DESC

	FieldSlice []string
	OrderSlice []string

	orderString string
}

func NewOrderBy(c *gin.Context) *OrderBy {
	return &OrderBy{Context: c}
}

func (p *OrderBy) Parse() error {
	if p.Fields = p.GetString("fields"); len(p.Fields) == 0 {
		return nil
	}
	p.Order = p.GetString("order")
	if len(p.Order) == 0 {
		p.orderString = p.Fields + defaultOrder
		return nil
	}
	p.OrderSlice = strings.Split(p.Order, separator)
	lenOrderSlice := len(p.OrderSlice)
	if lenOrderSlice == 1 {
		p.orderString = p.Fields + " " + p.Order
		return nil
	}
	p.FieldSlice = strings.Split(p.Fields, separator)
	if lenOrderSlice != len(p.FieldSlice) {
		return errors.New("fields or order error")
	}
	for idx, col := range p.FieldSlice {
		p.orderString += col + " " + p.OrderSlice[idx] + separator
	}
	p.orderString = strings.TrimRight(p.orderString, separator)
	return nil
}

func (p *OrderBy) String() string {
	return p.orderString
}
