package handler


import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mhmdryhn/cognito-auth-service/core/validation"
	"github.com/mhmdryhn/cognito-auth-service/schema"
	"github.com/mhmdryhn/cognito-auth-service/service/auth"
)


//  Validate incoming JSON data for a new user SignUp
func ValidateSignUpData(body []byte, signUpSchema *schema.SignUpSchema) (map[string]string, bool) {
	// Check if request body is well formatted JSON
	if err := json.Unmarshal(body, &signUpSchema); err != nil {
        return map[string]string {
			"jsonFormatError": "Invalid JSON data.",
		}, false
	}
	// Validate data based on `signUpSchema`
	validator := validation.NewValidator()
	err := validator.Struct(signUpSchema)
	if err == nil {
		return map[string]string {}, true
	} else {
		err := validation.ToCustomErrorMessage(err)
		return err, false
	}
}


// Get JSON request body as `[]byte` from context
func GetRequestBodyAsByteArray(ctx *gin.Context) ([]byte, error) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	return body, err
}


// Handler function used to SignUp a new user
func SignUpHandler(ctx *gin.Context) {
	body, _ := GetRequestBodyAsByteArray(ctx)
	var signUpSchema schema.SignUpSchema
	// Validate input data
	errorMessages, ok := ValidateSignUpData(body, &signUpSchema)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data": map[string]string {},
			"errors": errorMessages,
			"message": "Invalid JSON input data.",
			"statusCode": http.StatusBadRequest,
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
