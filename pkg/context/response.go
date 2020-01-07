package context

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code   int          `json:"code,omitempty" example:"400"`
	Msg    string       `json:"msg,omitempty" example:"status bad request"`
	Errors []*gin.Error `json:"errors,omitempty" example:"{array}"`
}

// type ResponseData struct {
// 	Data interface{} `json:"data,omitempty" example:"{}"`
// }

type ResponsePaging struct {
	Paging
	Records interface{} `json:"records" example:"{array}"`
}

type Paging struct {
	Page       int   `json:"page" example:"1"`
	PageSize   int   `json:"page_size" example:"5"`
	TotalCount int64 `json:"total_count" example:"100"`
}
