package handler


import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mhmdryhn/cognito-auth-service/schema"
	"github.com/mhmdryhn/cognito-auth-service/schema/validation"  // schemavalidation
	"github.com/mhmdryhn/cognito-auth-service/service/auth"
)


// Handler function used to SignUp a new user
func SignUpHandler(ctx *gin.Context) {
	body, _ := schemavalidation.GetRequestBodyAsByteArray(ctx)
	var signUpSchema schema.SignUpSchema
	// Validate input data
	errorMessages, ok := schemavalidation.ValidateSignUpData(body, &signUpSchema)
	if !ok {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"data": map[string]string {},
			"errors": errorMessages,
			"message": "Invalid JSON input data.",
			"statusCode": http.StatusUnprocessableEntity,
		})
		return
	} 
	// SignUp a new user here.
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
