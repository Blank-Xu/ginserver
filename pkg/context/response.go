package context

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code   int          `json:"code,omitempty" example:"4000"`
	Msg    string       `json:"msg,omitempty" example:"Invalid Params"`
	Errors []*gin.Error `json:"errors,omitempty" example:"[{\"error\":\"Invalid Params\",\"meta\":\"Invalid Params\"}]"`
}

// type ResponseData struct {
// 	Data interface{} `json:"data,omitempty"`
// }

type ResponsePaging struct {
	Paging  Paging      `json:"paging"`
	Records interface{} `json:"records"`
}

type Paging struct {
	Page       int   `json:"page" example:"1"`
	PageSize   int   `json:"page_size" example:"5"`
	TotalCount int64 `json:"total_count" example:"100"`
}
