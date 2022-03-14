package api

import (
	"net/http"

	"github.com/byteso/Xcloud/internal/cloud-client/service"
	"github.com/gin-gonic/gin"
)

func StorageEndpoint(c *gin.Context) {
	p := c.Param("path")

	switch p {
	case "storageInfo":
		storageInfo(c)
	}
}

func storageInfo(c *gin.Context) {
	account := c.MustGet("account").(string)

	response, err := service.GetStorageInfo(account)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  http.StatusText(http.StatusBadRequest),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": response,
		"msg":  http.StatusText(http.StatusOK),
	})
}
