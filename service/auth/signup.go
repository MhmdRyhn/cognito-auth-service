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
        UserAttributes: []*cognitoidp.AttributeType {
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

	_, err := client.AdminCreateUser(newUserInput)
	if err != nil {
		return "", err
	} else {
		return fmt.Sprintf("User with email %s created successfully.", username), err
	}
}


// Force change password on first login after admin creates the user
func ForceChangePassword(session string, username string, password string) (string, error) {
	appClientId := os.Getenv("APP_CLIENT_ID")
	challengeName := "NEW_PASSWORD_REQUIRED"
	respondToAuthChallengeInput := &cognitoidp.RespondToAuthChallengeInput {
		ClientId: aws.String(appClientId),
		ChallengeName: aws.String(challengeName),
		Session: aws.String(session),
		ChallengeResponses: map[string]*string {
			"USERNAME": aws.String(username),
			"NEW_PASSWORD": aws.String(password),
		},
	}

	_, err := client.RespondToAuthChallenge(respondToAuthChallengeInput)
	if err != nil {
		return "", err
	} else {
		return fmt.Sprintf("Password changed successfully."), err
	}
}


// Admin deletes a user from the User Pool
func AdminDeleteUser(username string) (string, error) {
	userPoolId := os.Getenv("USER_POOL")
	adminDeleteUserInput := &cognitoidp.AdminDeleteUserInput {
		UserPoolId: aws.String(userPoolId),
		Username: aws.String(username),
	}

	_, err := client.AdminDeleteUser(adminDeleteUserInput)
	if err != nil {
		return "", err
	} else {
		return fmt.Sprintf("User %s has been deleted successfully.", username), err
	}
}