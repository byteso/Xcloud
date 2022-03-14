package api

import (
	"net/http"

	"github.com/byteso/Xcloud/api/cloud-server/v1/types"
	"github.com/byteso/Xcloud/internal/cloud-server/service"
	"github.com/gin-gonic/gin"
)

func UserInfoEndpoint(c *gin.Context) {
	p := c.Param("path")

	switch p {
	case "userInfo":
		userInfo(c)
	}
}

func userInfo(c *gin.Context) {
	var request types.RequestUserInfo

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  http.StatusText(http.StatusBadRequest),
		})
		return
	}

	response, err := service.GetUserInfo(request)
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
