package route


import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mhmdryhn/cognito-auth-service/route/handler"
)


func RegisterRoutes(router *gin.Engine) {
	router.NoRoute(RouteNotFoundHandler)

	auth := router.Group("/auth")
	{
		// auth.POST("/create-user", auth.CreateUser)
		auth.POST("/signup", handler.SignUpHandler)
		auth.POST("/confirm-signup", handler.ConfirmSignUpHandler)
		auth.POST("/signin", handler.SignInHandler)
		auth.POST("/refresh-token-auth", handler.RefreshTokenAuthHandler)
		auth.POST("/forgot-password", handler.ForgotPasswordHandler)
		auth.POST("/confirm-forget-password", handler.ConfirmForgetPasswordHandler)
		// auth.POST("/change-passowrd", handler.ChangePassowrdHandler)

		// auth.POST("/force-change-password", handler.ForceChangePasswordHandler)
	}
}


func RouteNotFoundHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{
		"data": map[string]string {},
		"error": map[string]string {
			"RouteNotFoundOrMethodNotAllowed": "Requested route not found or method not allowed.",
		},
		"message": "Requested route not found or method not allowed.",
		"statusCode": http.StatusNotFound,
	})
}
