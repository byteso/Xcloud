package userinfo

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

func Get(c *gin.Context) {

}

func Update(c *gin.Context) {

}
