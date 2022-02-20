package api

import "github.com/gin-gonic/gin"

func UserInfoHandle(c *gin.Context) {
	path := c.Param("path")
	switch path {
	case "get":
		Get(c)
	case "update":
		Update(c)
	}
}

func GetInfo(c *gin.Context) {

}

func Update(c *gin.Context) {

}
