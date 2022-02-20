package api

import "github.com/gin-gonic/gin"

func SourceHandle(c *gin.Context) {
	path := c.Param("path")
	switch path {
	case "get":
		Get(c)
	case "upload":
		Upload(c)
	case "download":
		Download(c)
	case "delete":
		Delete(c)
	}
}

func Get(c *gin.Context) {

}

func Upload(c *gin.Context) {

}

func Download(c *gin.Context) {

}

func Delete(c *gin.Context) {

}
