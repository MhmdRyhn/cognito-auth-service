package route


import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mhmdryhn/cognito-auth-service/route/handler"
)


func RegisterRoutes(router *gin.Engine) {
	router.NoRoute(RouteNotFoundHandler)

	auth := router.Group("/auth")
	RegisterAuthRoutes(auth)
}


func RegisterAuthRoutes(authGroup *gin.RouterGroup) {
	{
		authGroup.POST("/signup", handler.SignUpHandler)
		authGroup.POST("/confirm-signup", handler.ConfirmSignUpHandler)
		authGroup.POST("/signin", handler.SignInHandler)
		authGroup.POST("/refresh-token-auth", handler.RefreshTokenAuthHandler)
		authGroup.POST("/forgot-password", handler.ForgotPasswordHandler)
		authGroup.POST("/confirm-forgot-password", handler.ConfirmForgetPasswordHandler)
	}
	authChangePassword := authGroup.Group("/")
	authChangePassword.Use(AuthenticationMiddeware())
	authChangePassword.POST("/change-password", handler.ChangePassword)
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
