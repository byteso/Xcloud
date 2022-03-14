package api

import (
	"fmt"
	"net/http"

	"github.com/byteso/Xcloud/internal/config"
	"github.com/gin-gonic/gin"
)

func ConnectGetEndpoint(c *gin.Context) {
	p := c.Param("path")
	fmt.Println(p)

	switch p {
	case "connectStatus":
		connectStatus(c)
	}
}

func connectStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": config.Config.CloudServer.Platform,
		"msg":  http.StatusText(http.StatusOK),
	})
}
