package api

import (
	"net/http"

	"github.com/byteso/Xcloud/api/cloud-client/v1/types"
	"github.com/byteso/Xcloud/internal/cloud-client/service"
	"github.com/gin-gonic/gin"
)

func UserInfoEndpoint(c *gin.Context) {
	path := c.Param("path")
	switch path {
	case "getInfo":
		getInfo(c)
	case "updateInfo":
		updateInfo(c)
	}
}

func getInfo(c *gin.Context) {
	account := c.MustGet("account").(string)

	response, err := service.GetInfo(account)
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

func updateInfo(c *gin.Context) {
	var request types.ResquestUpdateInfo

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  http.StatusText(http.StatusBadRequest),
		})
		return
	}

	account := c.MustGet("account").(string)

	err := service.UpdateInfo(account, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  http.StatusText(http.StatusBadRequest),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  http.StatusText(http.StatusOK),
	})
}
