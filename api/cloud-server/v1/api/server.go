package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/byteso/Xcloud/internal/cloud-server/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ServerEndpoint(c *gin.Context) {
	p := c.Param("path")

	switch p {
	case "serverInfo":
		serverInfo(c)
	}
}

func serverInfo(c *gin.Context) {
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	defer ws.Close()

	mt, message, err := ws.ReadMessage()
	if err != nil {
		return
	}

	fmt.Println(string(message))

	for {

		response, err := service.ServerInfo()
		/*
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
		*/

		r, err := json.Marshal(response)
		if err != nil {
			return
		}

		err = ws.WriteMessage(mt, []byte(r))
		if err != nil {
			return
		}
	}
}
