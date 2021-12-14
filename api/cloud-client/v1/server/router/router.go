package router

import (
	"github.com/byteso/Xcloud/api/cloud-client/v1/login"
	"github.com/byteso/Xcloud/api/cloud-client/v1/server/middleware"
	"github.com/byteso/Xcloud/api/cloud-client/v1/source"
	userinfo "github.com/byteso/Xcloud/api/cloud-client/v1/user-info"
	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.POST("/:path", login.LoginHandle)
	}

	auth := r.Group("/v1", middleware.AuthJwt())
	{
		auth.GET("/source/:path", source.SourceHandle)
		auth.GET("/info/:path", userinfo.UserInfoHandle)
	}
}
