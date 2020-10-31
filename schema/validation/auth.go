// This package is being used to validate JSON input. 
// For now, the function stucture is same, but it will 
// be converted into a common function for input 
// validation later.

package schemavalidation


import (
	"encoding/json"

	"github.com/mhmdryhn/cognito-auth-service/core/validation"
)


func ValidateJSONData(requestBody []byte, schemaStruct interface{}) (map[string]string, bool) {
	// Load JSON into a struct
	if err := json.Unmarshal(requestBody, &schemaStruct); err != nil {
		return map[string]string {
			"jsonFormatError": "Invalid JSON data.",
		}, false
	}
	// Validate data
	validator := validation.NewValidator()
	err := validator.Struct(schemaStruct)
	if err == nil {
		return map[string]string {}, true
	} else {
		err := validation.ToCustomErrorMessage(err)
		return err, false
	}
}
