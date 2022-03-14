package router

import (
	"github.com/byteso/Xcloud/api/cloud-server/v1/api"
	"github.com/byteso/Xcloud/api/cloud-server/v1/middleware"
	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.POST("/:path", api.LoginEndpoint)
		//v1.POST("/:path", api.InvitationEndpoint)
		v1.GET("/connect/:path", api.ConnectGetEndpoint)
		v1.GET("/serverInfo/:path", api.ServerEndpoint)
		v1.GET("/bucket/:path", api.BucketEndpoint)
	}

	auth := r.Group("/v1/auth", middleware.AuthJwt())
	{
		auth.POST("/invitationCode/:path", api.InvitationEndpoint)
		auth.GET("/serverInfo/:path", api.ServerEndpoint)
		auth.GET("/bucket/:path", api.BucketEndpoint)
	}

	r.Run("localhost:8081")
}
