package auth


import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	// "github.com/aws/aws-sdk-go/aws/awserr"
	// "github.com/aws/aws-sdk-go/aws/credentials"
	// "github.com/aws/aws-sdk-go/aws/session"
	cognitoidp "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	cognitoidpiface "github.com/aws/aws-sdk-go/service/cognitoidentityprovider/cognitoidentityprovideriface"
)


type CognitoAuth struct {
	Client cognitoidpiface.CognitoIdentityProviderAPI
	UserPoolId string
	AppClientId string

	// Available methods:
	//
	// func (self *CognitoAuth) SignUp(username string, password string) (string, error)
	// func (self *CognitoAuth) ConfirmSignUp(username string, confirmationCode string) (string, error)
	// func (self *Cognito) SignIn(username string, password string) (map[string]string, error)
}


func NewCognitoAuth() CognitoAuth {
	return CognitoAuth {
		Client: CognitoClient,
		UserPoolId: os.Getenv("USER_POOL_ID"),
		AppClientId: os.Getenv("APP_CLIENT_ID"),
	}
}


// Users Signup by themselves
func (self *CognitoAuth) SignUp(username string, password string) (map[string]string, error) {
	signUpInput := &cognitoidp.SignUpInput {
		ClientId: aws.String(self.AppClientId),
		Username: aws.String(username),
		Password: aws.String(password),
        UserAttributes: []*cognitoidp.AttributeType{
            {
                Name:  aws.String("email"),
                Value: aws.String(username),
			},
		},
	}

	response, err := self.Client.SignUp(signUpInput)
	if err != nil {
		return map[string]string {}, err
	} else {
		return map[string]string {
			"message": fmt.Sprintf(
				"User with email %s signed up successfully. Please check your email for confirmation code.", 
				username,
			),
			"cognitoUsername": *(response.UserSub),
		}, err
	}
}


// Confirm Signup by providing a confirmation email sent to the user email
func (self *CognitoAuth) ConfirmSignUp(username string, confirmationCode string) (string, error) {
	confirmSignupInput := &cognitoidp.ConfirmSignUpInput {
		ClientId: aws.String(self.AppClientId),
		Username: aws.String(username),
		ConfirmationCode: aws.String(confirmationCode),
	}

	_, err := self.Client.ConfirmSignUp(confirmSignupInput)
	if err != nil {
		return "", err
	} else {
		return fmt.Sprintf("User with email %s confirmed successfully.", username), err
	}
}


// User signin using `username` and `password`
func (self *CognitoAuth) SignIn(username string, password string) (map[string]string, error) {
	signinInput := &cognitoidp.InitiateAuthInput {
		ClientId : aws.String(self.AppClientId),
		AuthFlow: aws.String("USER_PASSWORD_AUTH"),
		AuthParameters: map[string]*string {
			"USERNAME": aws.String(username),
			"PASSWORD": aws.String(password),
		},
	}

	response, err := self.Client.InitiateAuth(signinInput)
	if err != nil {
		return map[string]string {}, err
	} else {
		if response.AuthenticationResult != nil {
			return map[string]string {
				"accessToken": *(response.AuthenticationResult.AccessToken),
				"idToken": *(response.AuthenticationResult.IdToken),
				"refreshToken": *(response.AuthenticationResult.RefreshToken),
			}, err
		} else {
			return map[string]string {
				"session": *(response.Session),
			}, err
		}
	}
}
