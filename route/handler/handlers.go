package handler


import (
	"net/http"
	// "os"

	"github.com/gin-gonic/gin"

	"github.com/mhmdryhn/cognito-auth-service/schema"
	"github.com/mhmdryhn/cognito-auth-service/schema/validation"  // schemavalidation
	"github.com/mhmdryhn/cognito-auth-service/service/auth"
)


var cognitoAuth auth.CognitoAuth = auth.NewCognitoAuth()


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
	response, awsError := cognitoAuth.SignUp(signUpSchema.Email, signUpSchema.Password)
	if awsError == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"data": map[string]string {"cognitoUsername": response["cognitoUsername"]},
			"errors": map[string]string {},
			"message": response["message"],
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
	// Validate input data
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
	response, awsError := cognitoAuth.ConfirmSignUp(
		confirmSignUpSchema.Email, 
		confirmSignUpSchema.ConfirmationCode,
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


// Handler function used to sign in a user
func SignInHandler(ctx *gin.Context) {
	body, _ := schemavalidation.GetRequestBodyAsByteArray(ctx)
	var signInSchema schema.SignInSchema
	// Validate input data
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


// Handler function used to get new tokens using Refresh token
func RefreshTokenAuthHandler(ctx *gin.Context) {
	body, _ := schemavalidation.GetRequestBodyAsByteArray(ctx)
	var refreshTokenAuthSchema schema.RefreshTokenAuthSchema
	// Validate input data
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
	// Get new tokens
	response, awsError := cognitoAuth.RefreshTokenAuth(refreshTokenAuthSchema.RefreshToken)
	if awsError == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"data": response,
			"errors": map[string]string {},
			"message": "Got new tokens successfully.",
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
func ForgotPasswordHandler(ctx *gin.Context) {
	body, _ := schemavalidation.GetRequestBodyAsByteArray(ctx)
	var forgetPasswordSchema schema.ForgetPasswordSchema
	// Validate input data
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
	// Send verification code to reset password
	response, awsError := cognitoAuth.ForgotPassword(forgetPasswordSchema.Email)
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
	// Validate input data
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
	// Reset password
	response, awsError := cognitoAuth.ConfirmForgotPassword(
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


// Change password staying signed in
func ChangePassword(ctx *gin.Context) {
	body, _ := schemavalidation.GetRequestBodyAsByteArray(ctx)
	var changePasswordSchema schema.ChangePasswordSchema
	// Validate input data
	errorMessages, ok := schemavalidation.ValidateJSONData(body, &changePasswordSchema)
	if !ok {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"data": map[string]string {},
			"errors": errorMessages,
			"message": "Invalid JSON input data.",
			"statusCode": http.StatusUnprocessableEntity,
		})
		return
	}
	// Change password
	response, awsError := cognitoAuth.ChangePassword(
		changePasswordSchema.CurrentPassword, 
		changePasswordSchema.NewPassword,
		changePasswordSchema.AccessToken,
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
