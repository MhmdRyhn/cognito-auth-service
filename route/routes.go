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
		auth.POST("/confirm-signup", handler.ConfirmSignUpHandler)
		auth.POST("/signin", handler.SignInHandler)
		// auth.POST("/refresh-token-auth", handler.RefreshTokenAuthHandler)
		auth.POST("/forget-password", handler.ForgetPasswordHandler)
		auth.POST("/confirm-forget-password", handler.ConfirmForgetPasswordHandler)
		// auth.POST("/change-passowrd", handler.ChangePassowrdHandler)

		// auth.POST("/force-change-password", handler.ForceChangePasswordHandler)
	}
}
