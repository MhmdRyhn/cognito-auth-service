package route


import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func AuthenticationMiddeware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": "From authentication middleware",
			"statusCode": http.StatusOK,
		})
	}
}
