package auth


import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	cognitoidp "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)


func awsSession() *session.Session {
	awsProfile := os.Getenv("AWS_PROFILE")
	awsRegion := os.Getenv("AWS_REGION")
	session, _ := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion),
		Credentials: credentials.NewSharedCredentials("", awsProfile),
	})
	return session
}

var client *cognitoidp.CognitoIdentityProvider = cognitoidp.New(awsSession())
