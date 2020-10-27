package handler


import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	// "net/http"

	"github.com/gin-gonic/gin"
	// validate "github.com/go-playground/validator/v10"

	"github.com/mhmdryhn/cognito-auth-service/core/validation"
	// "github.com/mhmdryhn/cognito-auth-service/service/auth"
	"github.com/mhmdryhn/cognito-auth-service/schema"
)


func GetRequestBodyAsByteArray(ctx *gin.Context) ([]byte, error) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	return body, err
}


func ValidateSignUpSchema(body []byte, signUpSchema *schema.SignUpSchema) {
	if err := json.Unmarshal(body, &signUpSchema); err != nil {
        panic(err)
	}
	
	// fmt.Println(fmt.Sprintf("Email: %s", signUpSchema.Email))
	// fmt.Println(fmt.Sprintf("Password: %s", signUpSchema.Password))
	// fmt.Println(fmt.Sprintf("ConfirmPassword: %s", signUpSchema.ConfirmPassword))

	validator := validation.NewValidator()

	// validator := validate.New()
	// validator.RegisterValidation("password", func(field validate.FieldLevel) bool{
	// 	return validation.PasswordValidator(field.Field().String())
	// })

	err := validator.Struct(signUpSchema)
	if err == nil {
		fmt.Println("Valid JSON")
	} else {
		// translator, _ := validation.NewTranslator()
		errs := validation.ToCustomErrorMessage(err)
		fmt.Println(errs)
		// for _, err := range err.(validate.ValidationErrors) {
		// 	// fmt.Println(err.Translate(translator))
		// 	fmt.Println(err.Namespace())
		// 	fmt.Println(err.Field())
		// 	fmt.Println(err.StructNamespace())
		// 	fmt.Println(err.StructField())
		// 	fmt.Println(err.Tag())
		// 	fmt.Println(err.ActualTag())
		// 	fmt.Println(err.Kind())
		// 	fmt.Println(err.Type())
		// 	fmt.Println(err.Value())
		// 	fmt.Println(err.Param())
		// 	// fmt.Println(err.JSON())
		// 	fmt.Println()
		// }
	}
}


func SignUpHandler(ctx *gin.Context) {
	// json := ctx.Request.Body
	// fmt.Println(json)
	// body, _ := ioutil.ReadAll(ctx.Request.Body)
	body, _ := GetRequestBodyAsByteArray(ctx)
    // fmt.Println(string(body))
    // fmt.Println(body)
	// fmt.Println("Request received")
	// var data SignUpSchema
	// if err := ctx.ShouldBindJSON(&data); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	// fmt.Printf("Email: %s, Password: %s\n", data.Email, data.Password)

	var signUpSchema schema.SignUpSchema
	ValidateSignUpSchema(body, &signUpSchema)
    // if err := json.Unmarshal(body, &obj); err != nil {
    //     panic(err)
    // }

	// fmt.Println(fmt.Sprintf("Email: %s", obj.Email))
	// fmt.Println(fmt.Sprintf("Password: %s", obj.Password))
	// fmt.Println(fmt.Sprintf("ConfirmPassword: %s", obj.ConfirmPassword))
	// fmt.Println(fmt.Sprintf("Check.One: %s", obj.Check[0].One))
	// fmt.Println(fmt.Sprintf("Check.Two: %s", obj.Check[0].Two))
	// fmt.Println(fmt.Sprintf("Check: %s", obj.Check))
	
	ctx.JSON(200, gin.H{"status": "Valid credentials."})
}
