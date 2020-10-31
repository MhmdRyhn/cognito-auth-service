package schemavalidation


import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"

	"github.com/mhmdryhn/cognito-auth-service/core/validation"
)


// Get JSON request body as `[]byte` from context
func GetRequestBodyAsByteArray(ctx *gin.Context) ([]byte, error) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	return body, err
}


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
