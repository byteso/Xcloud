package api

import (
	"fmt"
	"net/http"

	"github.com/byteso/Xcloud/api/cloud-client/v1/types"
	"github.com/byteso/Xcloud/internal/cloud-client/service"
	"github.com/gin-gonic/gin"
)

//
func LoginEndpoint(c *gin.Context) {
	p := c.Param("path")
	switch p {
	case "verifyInvitation":
		verifyInvitation(c)
	case "sign":
		sign(c)
	case "login":
		login(c)
	}
}

func verifyInvitation(c *gin.Context) {
	var request types.RequestInvitation

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  http.StatusText(http.StatusBadRequest),
		})
		return
	}

	response, err := service.VerifyInvitation(request)
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

func sign(c *gin.Context) {
	var request types.RequestSign

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  http.StatusText(http.StatusBadRequest),
		})
		return
	}

	err := service.Sign(request)
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

func login(c *gin.Context) {
	var request types.RequestLogin

	// bad resquest
	if c.BindJSON(&request) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  http.StatusText(http.StatusBadRequest),
		})
		return
	}

	token, err := service.Login(request)
	// internal server error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	// ok
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": token,
		"msg":  http.StatusText(http.StatusOK),
	})

}
