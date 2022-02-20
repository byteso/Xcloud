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
		v1.POST("/:path", api.InvitationHandle)
		v1.GET("/:path", api.ConnectGetHandle)
	}

	auth := r.Group("/v1/auth", middleware.AuthJwt())
	{
		auth.POST("/:path", api.InvitationHandle)
	}

	r.Run()
}
