package api

import "github.com/gin-gonic/gin"

func LoginHandle(c *gin.Context) {
	path := c.Param("path")

	switch path {
	case "login":
		login(c)
	}
}

func login(c *gin.Context) {

}
