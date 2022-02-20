package api

import (
	"fmt"
	"net/http"

	"github.com/byteso/Xcloud/internal/config"
	"github.com/gin-gonic/gin"
)

func ConnectGetHandle(c *gin.Context) {
	p := c.Param("path")
	fmt.Println(p)

	switch p {
	case "connect":
		connect(c)
	}
}

func connect(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": config.Config.CloudServer.Platform,
		"msg":  http.StatusText(http.StatusOK),
	})
}
