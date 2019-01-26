package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func PostLogin(c *gin.Context) {

}
