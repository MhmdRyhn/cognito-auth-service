package route


import (
	"github.com/gin-gonic/gin"

	"github.com/mhmdryhn/cognito-auth-service/route/handler"
)


func RegisterRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		// auth.POST("/create-user", auth.CreateUser)
		auth.POST("/signup", handler.SignUpHandler)
		// auth.POST("/signin", auth.SignIn)
		// auth.POST("/force-change-password", auth.ForceChangePassword)
		// auth.POST("/forget-password", auth.ForgetPassword)
		// auth.POST("/confirm-forget-password", auth.ConfirmForgetPassword)
		// auth.POST("/refresh-token-auth", auth.RefreshTokenAuth)
		// auth.POST("/change-passowrd", auth.ChangePassowrd)
	}
}
