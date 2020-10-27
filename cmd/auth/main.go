package main


import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/mhmdryhn/cognito-auth-service/route"
	// "github.com/mhmdryhn/cognito-auth-service/service/auth"
)


func printMap(response map[string]string) {
	for key, value := range response {
		fmt.Println(fmt.Sprintf("%s: %s", key, value))
		fmt.Println()
	}
}

func main()  {
	router := gin.Default()
	route.RegisterRoutes(router)
	router.Run("localhost:8080")

	// username := ""
	// password := ""
	// newPassword := ""
	// confirmationCode := ""
	// session := ""
	// accessToken := ""
	// refreshToken := ""
	
	// groupName := ""
	// description := ""
	// var precedence int64 = 0

	// response, err := auth.SignUp(username, password)
	// response, err := auth.ConfirmSignUp(username, confirmationCode)
	// response, err := auth.AdminCreateUser(username, password)
	// response, err := auth.SignIn(username, password)
	// response, err := auth.ForceChangePassword(session, username, password)
	// response, err := auth.RefreshTokenAuth(refreshToken)
	// response, err := auth.ForgetPassword(username)
	// response, err := auth.ConfirmForgetPassword(username, confirmationCode, password)
	// response, err := auth.ChangePassword(password, newPassword, accessToken)
	// response, err := auth.ResendConfirmationCode(username)

	// response, err := auth.CreateUserGroup(groupName, description, precedence)
	// response, err := auth.AdminAddUserToGroup(username, groupName)

	// response, err := auth.AdminDeleteUser(username)
	// response, err := auth.DeleteUserGroup(groupName)
	// response, err := auth.AdminRemoveUserFromGroup(username, groupName)

	// if err != nil {
	// 	code, msg := auth.CognitoErrorDetails(err)
	// 	fmt.Println(fmt.Sprintf("ErrorCode: %s; ErrorMessage: %s", code, msg))
	// } else {
	// 	fmt.Println(response)
	// 	// printMap(response)
	// }
}
