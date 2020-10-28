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
