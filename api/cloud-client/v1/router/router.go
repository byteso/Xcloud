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
		v1.POST("/:path", api.LoginEndpoint)
	}

	auth := r.Group("/v1/auth", middleware.AuthJwt())
	{
		// source
		auth.GET("/source/:path", api.SourceEndpoint)
		auth.POST("/source/:path", api.SourceEndpoint)
		auth.DELETE("/source/:path", api.SourceEndpoint)

		// user info
		auth.GET("/info/:path", api.UserInfoEndpoint)
		auth.POST("/info/:path", api.UserInfoEndpoint)

		// storage info
		auth.GET("storage/:path", api.StorageEndpoint)

		// file info
		auth.GET("/fileInfo/:path", api.FileInfoHandle)
	}

	r.Run("localhost:8080")
}
