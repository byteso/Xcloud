package router

import (
	"github.com/byteso/Xcloud/api/cloud-server/v1/api"
	"github.com/byteso/Xcloud/api/cloud-server/v1/middleware"
	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()

	test := r.Group("v1/test")
	{
		test.POST("/:path", api.InvitationHandle)
	}

	v1 := r.Group("/V1")
	{
		v1.POST("/:path", api.InvitationHandle)
	}

	auth := r.Group("/v1", middleware.AuthJwt())
	{
		auth.POST("/:path", api.InvitationHandle)
	}

	r.Run()
}
