package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wallet/auth"
)

func UserLoaderMiddleware() gin.HandlerFunc {
	return func (ctx *gin.Context) {
		err := auth.TokenValid(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
