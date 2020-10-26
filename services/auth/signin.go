package auth


import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	cognitoidp "github.com/aws/aws-sdk-go/service/cognitoidentityprovider" 
)


// User signin using `username` and `password`
func SignIn(username string, password string) (map[string]string, error) {
	appClientId := os.Getenv("APP_CLIENT_ID")
	signinInput := &cognitoidp.InitiateAuthInput {
		ClientId : aws.String(appClientId),
		AuthFlow: aws.String("USER_PASSWORD_AUTH"),
		AuthParameters: map[string]*string {
			"USERNAME": aws.String(username),
			"PASSWORD": aws.String(password),
		},
	}

	response, err := client.InitiateAuth(signinInput)
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
func RefreshTokenAuth(refreshToken string) (map[string]string, error) {
	appClientId := os.Getenv("APP_CLIENT_ID")
	refreshTokenAuthInput := &cognitoidp.InitiateAuthInput {
		ClientId : aws.String(appClientId),
		AuthFlow: aws.String("REFRESH_TOKEN_AUTH"),
		AuthParameters: map[string]*string {
			"REFRESH_TOKEN": aws.String(refreshToken),
		},
	}

	response, err := client.InitiateAuth(refreshTokenAuthInput)
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
