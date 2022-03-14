package api

import (
	"fmt"
	"net/http"

	"github.com/byteso/Xcloud/api/cloud-client/v1/types"
	"github.com/byteso/Xcloud/internal/cloud-client/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func SourceEndpoint(c *gin.Context) {
	path := c.Param("path")
	switch path {
	case "getSource":
		getSource(c)
	case "uploadSource":
		uploadSource(c)
	case "downloadSource":
		downloadSource(c)
	case "deleteSource":
		deleteSource(c)
	}
}

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func getSource(c *gin.Context) {
	bucketName := c.MustGet("bucketName").(string)
	id := c.Query("id")

	fmt.Println(id)

	response, err := service.GetSource(bucketName, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	c.DataFromReader(http.StatusOK, response.ContentLength, response.ContentType, response.Reader, response.ExtraHeaders)
}

func uploadSource(c *gin.Context) {
	var request types.RequestUploadSource

	bucketName := c.MustGet("bucketName").(string)

	form, _ := c.MultipartForm()
	files := form.File["upload[]"]

	err := service.UploadSource(bucketName, request.FolderName, files)
	fmt.Println(err)
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

func downloadSource(c *gin.Context) {
	var request types.RequestDownloadSource

	bucketName := c.MustGet("bucketName").(string)

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  http.StatusText(http.StatusBadRequest),
		})
		return
	}

	response, err := service.DownloadSource(bucketName, request.FolderName, request.Key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	c.Data(http.StatusOK, response.ContentType, response.Data)
}

func deleteSource(c *gin.Context) {
	var request types.RequestDeleteSource

	bucketName := c.MustGet("bucketName").(string)
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  http.StatusText(http.StatusBadRequest),
		})
		return
	}

	err := service.DeleteSource(bucketName, request.FolderName, request.Key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"code": http.StatusInternalServerError,
		"msg":  http.StatusText(http.StatusInternalServerError),
	})
}
