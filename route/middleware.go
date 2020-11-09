package route


import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/mhmdryhn/cognito-auth-service/core/utils"
)


func AuthenticationMiddeware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken := ctx.Request.Header.Get("Authorization")
		token := strings.Replace(accessToken, "Bearer ", "", 1)
		val := utils.JWTValidator{}
		if claims, err := val.Validate(token); err == nil {
			ctx.Set(
				"user", 
				map[string]interface{} {
					"username": claims["sub"], 
					// "group": claims["groups"],
				},
			)
			ctx.Next()
		} else {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"data": map[string]string {},
				"errors": map[string]string {"TokenError": err.Error()},
				"message": err.Error(),
				"statusCode": http.StatusUnauthorized,
			})
		}
	}
}
