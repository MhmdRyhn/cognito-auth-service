package schemavalidation


import (
	"encoding/json"

	"github.com/mhmdryhn/cognito-auth-service/core/validation"
	"github.com/mhmdryhn/cognito-auth-service/schema"
)


//  Validate incoming JSON data for a new user SignUp
func ValidateSignUpData(body []byte, signUpSchema *schema.SignUpSchema) (map[string]string, bool) {
	// Check if request body is well formatted JSON
	if err := json.Unmarshal(body, &signUpSchema); err != nil {
        return map[string]string {
			"jsonFormatError": "Invalid JSON data.",
		}, false
	}
	// Validate data based on `SignUpSchema`
	validator := validation.NewValidator()
	err := validator.Struct(signUpSchema)
	if err == nil {
		return map[string]string {}, true
	} else {
		err := validation.ToCustomErrorMessage(err)
		return err, false
	}
}


//  Validate incoming JSON data for a new user confirm SignUp
func ValidateConfirmSignUpData(body []byte, confirmSignUpSchema *schema.ConfirmSignUpSchema) (map[string]string, bool) {
	// Check if request body is well formatted JSON
	if err := json.Unmarshal(body, &confirmSignUpSchema); err != nil {
        return map[string]string {
			"jsonFormatError": "Invalid JSON data.",
		}, false
	}
	// Validate data based on `ConfirmSignUpSchema`
	validator := validation.NewValidator()
	err := validator.Struct(confirmSignUpSchema)
	if err == nil {
		return map[string]string {}, true
	} else {
		err := validation.ToCustomErrorMessage(err)
		return err, false
	}
}


//  Validate incoming JSON data for user signin
func ValidateSignInData(body []byte, signInSchema *schema.SignInSchema) (map[string]string, bool) {
	// Check if request body is well formatted JSON
	if err := json.Unmarshal(body, &signInSchema); err != nil {
        return map[string]string {
			"jsonFormatError": "Invalid JSON data.",
		}, false
	}
	// Validate data based on `SignInSchema`
	validator := validation.NewValidator()
	err := validator.Struct(signInSchema)
	if err == nil {
		return map[string]string {}, true
	} else {
		err := validation.ToCustomErrorMessage(err)
		return err, false
	}
}

