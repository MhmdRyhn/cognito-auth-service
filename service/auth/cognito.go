package auth


import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	cognitoidp "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)


// Create and then return an AWS session
func awsSession() *session.Session {
	awsProfile := os.Getenv("AWS_PROFILE")
	awsRegion := os.Getenv("AWS_REGION")
	session, _ := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion),
		Credentials: credentials.NewSharedCredentials("", awsProfile),
	})
	return session
}


// Cognito client
var CognitoClient *cognitoidp.CognitoIdentityProvider = cognitoidp.New(awsSession())


// Cognito client
var client *cognitoidp.CognitoIdentityProvider = cognitoidp.New(awsSession())


// Extracts `error code` and `error message` from the `awserr`
func CognitoErrorDetails(err error) (string, string) {
	return err.(awserr.Error).Code(), err.(awserr.Error).Message()
}


// Resend a verification code to users' `email` to reset password
func ResendConfirmationCode(username string) (string, error) {
	appClientId := os.Getenv("APP_CLIENT_ID")

	resendConfirmationCodeInput := &cognitoidp.ResendConfirmationCodeInput {
		ClientId: aws.String(appClientId),
		Username: aws.String(username),
    }

	_, err := client.ResendConfirmationCode(resendConfirmationCodeInput)
	if err != nil {
		return "", err
	} else {
		return fmt.Sprintf(
			"A verification code will be resent to email %s if a user exists with this email.", username,
		), err
	}
}
