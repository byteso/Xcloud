package api

import (
	"net/http"

	"github.com/byteso/Xcloud/api/cloud-server/v1/types"
	"github.com/byteso/Xcloud/internal/cloud-server/service"
	"github.com/gin-gonic/gin"
)

func InvitationEndpoint(c *gin.Context) {
	p := c.Param("path")
	switch p {
	case "createInvitation":
		createInvitation(c)
	case "invitationInfo":
		invitationInfo(c)
	}
}

func createInvitation(c *gin.Context) {
	var request types.RequestInvitation
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  http.StatusText(http.StatusBadRequest),
		})
		return
	}

	if err := service.CreateInvitation(request); err != nil {
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

func invitationInfo(c *gin.Context) {
	var request types.RequestInvitationInfo

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  http.StatusText(http.StatusBadRequest),
		})
		return
	}

	response, err := service.GetInvitationInfo(request)
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
