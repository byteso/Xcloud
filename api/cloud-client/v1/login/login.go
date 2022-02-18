package login

import (
	"net/http"

	typeLogin "github.com/byteso/Xcloud/api/cloud-client/v1/types/login"
	internalLogin "github.com/byteso/Xcloud/internal/cloud-client/service/login"
	"github.com/gin-gonic/gin"
)

//
func LoginHandle(c *gin.Context) {
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

}

func sign(c *gin.Context) {

}

func login(c *gin.Context) {
	var request typeLogin.RequestLogin

	// bad resquest
	if c.BindJSON(&request) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  http.StatusText(http.StatusBadRequest),
		})
		return
	}

	token, err := internalLogin.Login(request)
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
