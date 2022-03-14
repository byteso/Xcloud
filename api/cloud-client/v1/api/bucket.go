package api

import (
	"net/http"

	"github.com/byteso/Xcloud/api/cloud-client/v1/types"
	"github.com/byteso/Xcloud/internal/cloud-client/service"
	"github.com/gin-gonic/gin"
)

func BucketEndpoint(c *gin.Context) {
	p := c.Param("path")

	switch p {
	case "buckets":
		buckets(c)
	}
}

func buckets(c *gin.Context) {
	var request types.RequestCreateBucket

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  http.StatusText(http.StatusBadRequest),
		})
		return
	}

	err := service.CreateBucket(request.BucketName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  http.StatusText(http.StatusOK),
	})
}
