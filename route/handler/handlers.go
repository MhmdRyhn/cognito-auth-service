package handler


import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/mhmdryhn/cognito-auth-service/schema"
	"github.com/mhmdryhn/cognito-auth-service/schema/validation"  // schemavalidation
	"github.com/mhmdryhn/cognito-auth-service/service/auth"
)


var cognitoAuth auth.CognitoAuth = auth.CognitoAuth {
	Client: auth.CognitoClient,
	UserPoolId: os.Getenv("USER_POOL"),
	AppClientId: os.Getenv("APP_CLIENT_ID"),
}


// Handler function used to SignUp a new user
func SignUpHandler(ctx *gin.Context) {
	body, _ := schemavalidation.GetRequestBodyAsByteArray(ctx)
	var signUpSchema schema.SignUpSchema
	// Validate input data
	errorMessages, ok := schemavalidation.ValidateJSONData(body, &signUpSchema)
	if !ok {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"data": map[string]string {},
			"errors": errorMessages,
			"message": "Invalid JSON input data.",
			"statusCode": http.StatusUnprocessableEntity,
		})
		return
	} 
	// SignUp a new user
	response, awsError := auth.SignUp(signUpSchema.Email, signUpSchema.Password)
	if awsError == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"data": map[string]string {},
			"errors": map[string]string {},
			"message": response,
			"statusCode": http.StatusOK,
		})
	} else {
		errorCode, errorMessage := auth.CognitoErrorDetails(awsError)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data": map[string]string {},
			"errors": map[string]string {errorCode: errorMessage},
			"message": errorMessage,
			"statusCode": http.StatusBadRequest,
		})
	}
}


// Handler function used to confirm SignUp of a new user
func ConfirmSignUpHandler(ctx *gin.Context) {
	body, _ := schemavalidation.GetRequestBodyAsByteArray(ctx)
	var confirmSignUpSchema schema.ConfirmSignUpSchema
	// Validate confirm signup input data
	errorMessages, ok := schemavalidation.ValidateJSONData(body, &confirmSignUpSchema)
	if !ok {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"data": map[string]string {},
			"errors": errorMessages,
			"message": "Invalid JSON input data.",
			"statusCode": http.StatusUnprocessableEntity,
		})
		return
	} 
	// Confirm SignUp a new user
	response, awsError := auth.ConfirmSignUp(confirmSignUpSchema.Email, confirmSignUpSchema.ConfirmationCode)
	if awsError == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"data": map[string]string {},
			"errors": map[string]string {},
			"message": response,
			"statusCode": http.StatusOK,
		})
	} else {
		errorCode, errorMessage := auth.CognitoErrorDetails(awsError)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data": map[string]string {},
			"errors": map[string]string {errorCode: errorMessage},
			"message": errorMessage,
			"statusCode": http.StatusBadRequest,
		})
	}
}


// Handler function used to sign in a user
func SignInHandler(ctx *gin.Context) {
	body, _ := schemavalidation.GetRequestBodyAsByteArray(ctx)
	var signInSchema schema.SignInSchema
	// Validate signin input data
	errorMessages, ok := schemavalidation.ValidateJSONData(body, &signInSchema)
	if !ok {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"data": map[string]string {},
			"errors": errorMessages,
			"message": "Invalid JSON input data.",
			"statusCode": http.StatusUnprocessableEntity,
		})
		return
	}
	// Signin user
	response, awsError := cognitoAuth.SignIn(signInSchema.Email, signInSchema.Password)
	if awsError == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"data": response,
			"errors": map[string]string {},
			"message": "Signed in succssfully.",
			"statusCode": http.StatusOK,
		})
	} else {
		errorCode, errorMessage := auth.CognitoErrorDetails(awsError)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data": map[string]string {},
			"errors": map[string]string {errorCode: errorMessage},
			"message": errorMessage,
			"statusCode": http.StatusBadRequest,
		})
	}
}


// Handler function used to get new tokej using Refresh token
func RefreshTokenAuthHandler(ctx *gin.Context) {
	body, _ := schemavalidation.GetRequestBodyAsByteArray(ctx)
	var refreshTokenAuthSchema schema.RefreshTokenAuthSchema
	// Validate signin input data
	errorMessages, ok := schemavalidation.ValidateJSONData(body, &refreshTokenAuthSchema)
	if !ok {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"data": map[string]string {},
			"errors": errorMessages,
			"message": "Invalid JSON input data.",
			"statusCode": http.StatusUnprocessableEntity,
		})
		return
	}
	// Signin user
	response, awsError := auth.RefreshTokenAuth(refreshTokenAuthSchema.RefreshToken)
	if awsError == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"data": map[string]string {},
			"errors": map[string]string {},
			"message": response,
			"statusCode": http.StatusOK,
		})
	} else {
		errorCode, errorMessage := auth.CognitoErrorDetails(awsError)
		ctx.JSON(http.StatusBadRequest, gin.H {
			"data": map[string]string {},
			"errors": map[string]string {errorCode: errorMessage},
			"message": errorMessage,
			"statusCode": http.StatusBadRequest,
		})
	}
}


// Handler function used to get code for reset password
func ForgetPasswordHandler(ctx *gin.Context) {
	body, _ := schemavalidation.GetRequestBodyAsByteArray(ctx)
	var forgetPasswordSchema schema.ForgetPasswordSchema
	// Validate signin input data
	errorMessages, ok := schemavalidation.ValidateJSONData(body, &forgetPasswordSchema)
	if !ok {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"data": map[string]string {},
			"errors": errorMessages,
			"message": "Invalid JSON input data.",
			"statusCode": http.StatusUnprocessableEntity,
		})
		return
	}
	// Signin user
	response, awsError := auth.ForgetPassword(forgetPasswordSchema.Email)
	if awsError == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"data": map[string]string {},
			"errors": map[string]string {},
			"message": response,
			"statusCode": http.StatusOK,
		})
	} else {
		errorCode, errorMessage := auth.CognitoErrorDetails(awsError)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data": map[string]string {},
			"errors": map[string]string {errorCode: errorMessage},
			"message": errorMessage,
			"statusCode": http.StatusBadRequest,
		})
	}
}


// Handler function used to reset password
func ConfirmForgetPasswordHandler(ctx *gin.Context) {
	body, _ := schemavalidation.GetRequestBodyAsByteArray(ctx)
	var confirmForgetPasswordSchema schema.ConfirmForgetPasswordSchema
	// Validate signin input data
	errorMessages, ok := schemavalidation.ValidateJSONData(body, &confirmForgetPasswordSchema)
	if !ok {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"data": map[string]string {},
			"errors": errorMessages,
			"message": "Invalid JSON input data.",
			"statusCode": http.StatusUnprocessableEntity,
		})
		return
	}
	// Signin user
	response, awsError := auth.ConfirmForgetPassword(
		confirmForgetPasswordSchema.Email, 
		confirmForgetPasswordSchema.ConfirmationCode,
		confirmForgetPasswordSchema.Password,
	)
	if awsError == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"data": map[string]string {},
			"errors": map[string]string {},
			"message": response,
			"statusCode": http.StatusOK,
		})
	} else {
		errorCode, errorMessage := auth.CognitoErrorDetails(awsError)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data": map[string]string {},
			"errors": map[string]string {errorCode: errorMessage},
			"message": errorMessage,
			"statusCode": http.StatusBadRequest,
		})
	}
}
