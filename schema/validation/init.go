package schemavalidation


import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
)


// Get JSON request body as `[]byte` from context
func GetRequestBodyAsByteArray(ctx *gin.Context) ([]byte, error) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	return body, err
}
