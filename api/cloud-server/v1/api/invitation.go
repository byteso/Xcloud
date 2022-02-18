package api

import (
	"net/http"

	"github.com/byteso/Xcloud/api/cloud-server/v1/types"
	"github.com/byteso/Xcloud/internal/cloud-server/service"
	"github.com/gin-gonic/gin"
)

func InvitationHandle(c *gin.Context) {
	p := c.Param("path")
	switch p {
	case "invitation":
		invitation(c)
	case "invitationInfo":
		invitationInfo(c)
	}
}

func invitation(c *gin.Context) {
	var request types.RequestInvitation
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  http.StatusText(http.StatusBadRequest),
		})
		return
	}

	if err := service.AddInvitation(request); err != nil {
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

}
