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
	// func (self *CognitoAuth) SignUp(username string, password string) (map[string]string, error)
	// func (self *CognitoAuth) ConfirmSignUp(username string, confirmationCode string) (string, error)
	// func (self *CognitoAuth) SignIn(username string, password string) (map[string]string, error)
	// func (self *CognitoAuth) RefreshTokenAuth(refreshToken string) (map[string]string, error)
	// func (self *CognitoAuth) ForgetPassword(username string) (string, error)
	// func (self *CognitoAuth) ConfirmForgetPassword(username string, confirmationCode string, password string) (string, error)
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


// Get new `accessToken` and `idToken` using the `refreshToken`
func (self *CognitoAuth) RefreshTokenAuth(refreshToken string) (map[string]string, error) {
	refreshTokenAuthInput := &cognitoidp.InitiateAuthInput {
		ClientId : aws.String(self.AppClientId),
		AuthFlow: aws.String("REFRESH_TOKEN_AUTH"),
		AuthParameters: map[string]*string {
			"REFRESH_TOKEN": aws.String(refreshToken),
		},
	}

	response, err := self.Client.InitiateAuth(refreshTokenAuthInput)
	if err != nil {
		return map[string]string {}, err
	} else {
		return map[string]string {
			"accessToken": *(response.AuthenticationResult.AccessToken),
			"idToken": *(response.AuthenticationResult.IdToken),
			"refreshToken": refreshToken,
		}, err
	}
}


// Send a verification code to users' `email` to reset password
func (self *CognitoAuth) ForgotPassword(username string) (string, error) {
	forgetPasswordInput := &cognitoidp.ForgotPasswordInput {
		ClientId: aws.String(self.AppClientId),
		Username: aws.String(username),
    }

	_, err := self.Client.ForgotPassword(forgetPasswordInput)
	if err != nil {
		return "", err
	} else {
		return fmt.Sprintf(
			"A verification code will be sent to email %s if a user exists with this email.", username,
		), err
	}
}


// Users can reset password using the verification code sent to their email
func (self *CognitoAuth) ConfirmForgotPassword(username string, confirmationCode string, password string) (string, error) {
	confirmForgetPasswordInput := &cognitoidp.ConfirmForgotPasswordInput {
		ClientId: aws.String(self.AppClientId),
		Username: aws.String(username),
		ConfirmationCode: aws.String(confirmationCode),
		Password: aws.String(password),
    }

	_, err := self.Client.ConfirmForgotPassword(confirmForgetPasswordInput)
	if err != nil {
		return "", err
	} else {
		return fmt.Sprintf("New password has been set successfully."), err
	}
}


// Change password staying signed in
func (self *CognitoAuth) ChangePassword(currentPassword string, newPassword string, accessToken string) (string, error) {
	changePasswordInput := &cognitoidp.ChangePasswordInput {
		PreviousPassword: aws.String(currentPassword),
		ProposedPassword: aws.String(newPassword),
		AccessToken: aws.String(accessToken),
    }

	_, err := self.Client.ChangePassword(changePasswordInput)
	if err != nil {
		return "", err
	} else {
		return fmt.Sprintf("Password changed successfully."), err
	}
}
