package auth


import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	cognitoidp "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)


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
