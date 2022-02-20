package middleware

import (
	"fmt"

	"github.com/byteso/Xcloud/internal/auth"
	"github.com/gin-gonic/gin"
)

func AuthJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		const bearerSchema = "Bearer"
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(bearerSchema)+1:]
		fmt.Println(tokenString)
		auth.ParseToken(tokenString, "client")

	}
}
