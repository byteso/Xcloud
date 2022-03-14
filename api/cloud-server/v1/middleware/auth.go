package middleware

import (
	"fmt"
	"net/http"

	"github.com/byteso/Xcloud/internal/auth"
	"github.com/gin-gonic/gin"
)

func AuthJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		const bearerSchema = "Bearer"
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  http.StatusText(http.StatusBadRequest),
			})
			c.Abort()
			return
		}

		tokenString := authHeader[len(bearerSchema)+1:]
		fmt.Println(tokenString)

		if _, ok := auth.ParseToken(tokenString, "server"); !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  http.StatusText(http.StatusBadRequest),
			})
			c.Abort()
			return
		}
	}
}
