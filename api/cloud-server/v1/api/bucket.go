package api

import (
	"net/http"

	"github.com/byteso/Xcloud/internal/cloud-server/service"
	"github.com/gin-gonic/gin"
)

func BucketEndpoint(c *gin.Context) {
	p := c.Param("path")

	switch p {
	case "listBuckets":
		listBuckets(c)
	}
}

func listBuckets(c *gin.Context) {
	response, err := service.ListBuckets()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": response,
		"msg":  http.StatusText(http.StatusOK),
	})
}
