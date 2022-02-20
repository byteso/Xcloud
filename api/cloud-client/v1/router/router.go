package router

import (
	"github.com/byteso/Xcloud/api/cloud-client/v1/api"
	"github.com/byteso/Xcloud/api/cloud-client/v1/middleware"
	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.POST("/:path", api.LoginHandle)
	}

	auth := r.Group("/v1/auth", middleware.AuthJwt())
	{
		auth.GET("/source/:path", api.SourceHandle)
		auth.GET("/info/:path", api.UserInfoHandle)
	}
	r.Run()
}
