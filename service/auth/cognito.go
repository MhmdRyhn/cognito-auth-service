package auth


import (
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


// TODO: Remove this after refactoring `func ResendConfirmationCode(username string) (string, error)`
// Cognito client
var client *cognitoidp.CognitoIdentityProvider = cognitoidp.New(awsSession())


// Extracts `error code` and `error message` from the `awserr`
func CognitoErrorDetails(err error) (string, string) {
	return err.(awserr.Error).Code(), err.(awserr.Error).Message()
}
