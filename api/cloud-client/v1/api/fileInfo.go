package api

import (
	"fmt"
	"net/http"

	"github.com/byteso/Xcloud/internal/cloud-client/service"
	"github.com/gin-gonic/gin"
)

func FileInfoHandle(c *gin.Context) {
	p := c.Param("path")

	switch p {
	case "items":
		items(c)
	case "getFileInfo":
		getFileInfo(c)
	}
}

func items(c *gin.Context) {
	bucketName := c.MustGet("bucketName").(string)

	response, err := service.Items(bucketName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusOK,
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

func getFileInfo(c *gin.Context) {
	bucketName := c.MustGet("bucketName").(string)
	id := c.DefaultQuery("id", "")

	fmt.Println(id)
	response, err := service.GetFileInfo(bucketName, id)
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
