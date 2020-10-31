// This package is being used to validate JSON input. 
// For now, the function stucture is same, but it will 
// be converted into a common function for input 
// validation later.

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


//  Validate refresh token auth JSON data
func ValidateRefreshTokenAuthData(body []byte, refreshTokenAuthSchema *schema.RefreshTokenAuthSchema) (map[string]string, bool) {
	// Check if request body is well formatted JSON
	if err := json.Unmarshal(body, &refreshTokenAuthSchema); err != nil {
        return map[string]string {
			"jsonFormatError": "Invalid JSON data.",
		}, false
	}
	// Validate data based on `RefreshTokenAuthSchema`
	validator := validation.NewValidator()
	err := validator.Struct(refreshTokenAuthSchema)
	if err == nil {
		return map[string]string {}, true
	} else {
		err := validation.ToCustomErrorMessage(err)
		return err, false
	}
}


//  Validate forget password JSON data
func ValidateForgetPasswordData(body []byte, forgetPasswordSchema *schema.ForgetPasswordSchema) (map[string]string, bool) {
	// Check if request body is well formatted JSON
	if err := json.Unmarshal(body, &forgetPasswordSchema); err != nil {
        return map[string]string {
			"jsonFormatError": "Invalid JSON data.",
		}, false
	}
	// Validate data based on `ForgetPasswordSchema`
	validator := validation.NewValidator()
	err := validator.Struct(forgetPasswordSchema)
	if err == nil {
		return map[string]string {}, true
	} else {
		err := validation.ToCustomErrorMessage(err)
		return err, false
	}
}


//  Validate confirm forget password JSON data
func ValidateConfirmForgetPasswordData(body []byte, confirmForgetPasswordSchema *schema.ConfirmForgetPasswordSchema) (map[string]string, bool) {
	// Check if request body is well formatted JSON
	if err := json.Unmarshal(body, &confirmForgetPasswordSchema); err != nil {
        return map[string]string {
			"jsonFormatError": "Invalid JSON data.",
		}, false
	}
	// Validate data based on `ConfirmForgetPasswordSchema`
	validator := validation.NewValidator()
	err := validator.Struct(confirmForgetPasswordSchema)
	if err == nil {
		return map[string]string {}, true
	} else {
		err := validation.ToCustomErrorMessage(err)
		return err, false
	}
}
