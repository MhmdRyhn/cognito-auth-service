package auth


import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	cognitoidp "github.com/aws/aws-sdk-go/service/cognitoidentityprovider" 
)


// Send a verification code to users' `email` to reset password
func ForgetPassword(username string) (string, error) {
	userPoolId := os.Getenv("APP_CLIENT_ID")

	forgetPasswordInput := &cognitoidp.ForgotPasswordInput {
		ClientId: aws.String(userPoolId),
		Username: aws.String(username),
    }

	_, err := client.ForgotPassword(forgetPasswordInput)
	if err != nil {
		return "", err
	} else {
		return fmt.Sprintf(
			"A verification code will be sent to email %s if a user exist with this email.", username,
		), err
	}
}


// Users can reset password using the verification code sent to their email
func ConfirmForgetPassword(username string, confirmationCode string, password string) (string, error) {
	userPoolId := os.Getenv("APP_CLIENT_ID")

	confirmForgetPasswordInput := &cognitoidp.ConfirmForgotPasswordInput {
		ClientId: aws.String(userPoolId),
		Username: aws.String(username),
		ConfirmationCode: aws.String(confirmationCode),
		Password: aws.String(password),
    }

	_, err := client.ConfirmForgotPassword(confirmForgetPasswordInput)
	if err != nil {
		return "", err
	} else {
		return fmt.Sprintf("New password has been set successfully."), err
	}
}
