package login

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//
func LoginHandle(c *gin.Context) {
	path := c.Param("path")
	switch path {
	case "sign":
		sign(c)
	case "login":
		login(c)
	}
}

func login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  http.StatusText(http.StatusOK),
	})

}

func sign(c *gin.Context) {

}
