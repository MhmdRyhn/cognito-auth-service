package auth


import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	cognitoidp "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)


// Users Signup by themselves
func SignUp(username string, password string) (string, error) {
	appClientId := os.Getenv("APP_CLIENT_ID")
	newUserInput := &cognitoidp.SignUpInput {
		ClientId: aws.String(appClientId),
		Username: aws.String(username),
		Password: aws.String(password),
        UserAttributes: []*cognitoidp.AttributeType{
            {
                Name:  aws.String("email"),
                Value: aws.String(username),
			},
		},
	}

	// client := cognitoidp.New(awsSession())
	_, err := client.SignUp(newUserInput)
	if err != nil {
		return "", err
	} else {
		return fmt.Sprintf("User with email %s signed up successfully.", username), err
	}
}


// Confirm Signup by providing a confirmation email sent to the user email
func ConfirmSignUp(username string, confirmationCode string) (string, error) {
	appClientId := os.Getenv("APP_CLIENT_ID")
	confirmSignupInput := &cognitoidp.ConfirmSignUpInput {
		ClientId: aws.String(appClientId),
		Username: aws.String(username),
		ConfirmationCode: aws.String(confirmationCode),
	}

	// client := cognitoidp.New(awsSession())
	_, err := client.ConfirmSignUp(confirmSignupInput)
	if err != nil {
		return "", err
	} else {
		return fmt.Sprintf("User with email %s confirmed successfully.", username), err
	}
}


// Admin creates user with a temporary password
func AdminCreateUser(username string, temporaryPassword string) (string, error) {
	userPoolId := os.Getenv("USER_POOL")

	newUserInput := &cognitoidp.AdminCreateUserInput {
		UserPoolId: aws.String(userPoolId),
		Username: aws.String(username),
		TemporaryPassword: aws.String(temporaryPassword),
        DesiredDeliveryMediums: []*string{
            aws.String("EMAIL"),
        },
        UserAttributes: []*cognitoidp.AttributeType{
            {
                Name:  aws.String("email"),
                Value: aws.String(username),
			},
			{
                Name:  aws.String("email_verified"),
                Value: aws.String("true"),
            },
		},
    }

	// client := cognitoidp.New(awsSession())
	_, err := client.AdminCreateUser(newUserInput)
	if err != nil {
		return "", err
	} else {
		return fmt.Sprintf("User with email %s created successfully.", username), err
	}
}
