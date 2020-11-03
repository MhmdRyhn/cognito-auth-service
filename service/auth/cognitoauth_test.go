package auth


import (
	"fmt"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	cognitoidp "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	cognitoidpiface "github.com/aws/aws-sdk-go/service/cognitoidentityprovider/cognitoidentityprovideriface"
)


// >>> Test `CognitoAuth.Signup` method
type mockSignUp struct {
	cognitoidpiface.CognitoIdentityProviderAPI
	Response cognitoidp.SignUpOutput
}


func (mock mockSignUp) SignUp(input *cognitoidp.SignUpInput) (*cognitoidp.SignUpOutput, error) {
	return &mock.Response, nil
}


func TestCognitoAuthSignUp(t *testing.T) {
	email := "user@email.com"
	cognitoUsername := "7a2361dc-1140-458b-88a9-2addb6c3d183"
	testCases := []struct {
		Response cognitoidp.SignUpOutput
		Expected map[string]string
	} {
		{
			Response: cognitoidp.SignUpOutput {
				UserSub: aws.String(cognitoUsername),
			},
			Expected: map[string]string {
				"message": fmt.Sprintf(
					"User with email %s signed up successfully. Please check your email for confirmation code.", 
					email,
				),
				"cognitoUsername": cognitoUsername,
			},
		},
	}

	for _, c := range testCases {
		cognitoAuth := CognitoAuth {
			Client: mockSignUp{Response: c.Response},
			UserPoolId: "mock-user-pool-id",
			AppClientId: "mock-app-client-id",
		}
		response, err := cognitoAuth.SignUp(email, "Password!23")
		if err != nil {
			t.Error("Error while signing up.")
		}
		equal := reflect.DeepEqual(response, c.Expected)
		if ! equal {
			t.Error("Signup failed.")
		}
	}
}


// >>> Test `CognitoAuth.ConfirmSignup` method
type mockConfirmSignUp struct {
	cognitoidpiface.CognitoIdentityProviderAPI
	Response cognitoidp.ConfirmSignUpOutput
}


func (mock mockConfirmSignUp) ConfirmSignUp(input *cognitoidp.ConfirmSignUpInput) (*cognitoidp.ConfirmSignUpOutput, error) {
	return &mock.Response, nil
}


func TestCognitoAuthConfirmSignUp(t *testing.T) {
	email := "user@email.com"
	testCases := []struct {
		Response cognitoidp.ConfirmSignUpOutput
		Expected string
	} {
		{
			Response: cognitoidp.ConfirmSignUpOutput {},
			Expected: fmt.Sprintf("User with email %s confirmed successfully.", email),
		},
	}

	for _, c := range testCases {
		cognitoAuth := CognitoAuth {
			Client: mockConfirmSignUp {Response: c.Response},
			UserPoolId: "mock-user-pool-id",
			AppClientId: "mock-app-client-id",
		}
		response, err := cognitoAuth.ConfirmSignUp("user@email.com", "Password!23")
		if err != nil {
			t.Error("Error while confirming sign up.")
		}
		if response != c.Expected {
			t.Error("Confirm signup failed.")
		}
	}
}


// >>> Test `CognitoAuth.SignIn` method
type mockSignIn struct {
	cognitoidpiface.CognitoIdentityProviderAPI
	Response cognitoidp.InitiateAuthOutput
}


func (mock mockSignIn) InitiateAuth(input *cognitoidp.InitiateAuthInput) (*cognitoidp.InitiateAuthOutput, error) {
	return &mock.Response, nil
}


func TestCognitoAuthSignIn(t *testing.T) {
	accessToken := "access-token"
	idToken := "id-token"
	refreshToken := "refresh-token"
	testCases := []struct {
		Response cognitoidp.InitiateAuthOutput
		Expected map[string]string
	} {
		{
			Response: cognitoidp.InitiateAuthOutput {
				AuthenticationResult: &cognitoidp.AuthenticationResultType {
					AccessToken: aws.String(accessToken),
					ExpiresIn: aws.Int64(3600),
					IdToken: aws.String(idToken),
					RefreshToken: aws.String(refreshToken),
					TokenType: aws.String("Access"),
				},
			},
			Expected: map[string]string {
				"accessToken": accessToken,
				"idToken": idToken,
				"refreshToken": refreshToken,
			},
		},
	}

	for _, c := range testCases {
		cognitoAuth := CognitoAuth {
			Client: mockSignIn{Response: c.Response},
			UserPoolId: "mock-user-pool-id",
			AppClientId: "mock-app-client-id",
		}
		response, err := cognitoAuth.SignIn("user@email.com", "Password!23")
		if err != nil {
			t.Error("Error while signing in.")
		}
		equal := reflect.DeepEqual(response, c.Expected)
		if ! equal {
			t.Error("Signin failed.")
		}
	}
}


// >>> Test `CognitoAuth.RefreshTokenAuth` method
type mockRefreshTokenAuth struct {
	cognitoidpiface.CognitoIdentityProviderAPI
	Response cognitoidp.InitiateAuthOutput
}


func (mock mockRefreshTokenAuth) InitiateAuth(input *cognitoidp.InitiateAuthInput) (*cognitoidp.InitiateAuthOutput, error) {
	return &mock.Response, nil
}


func TestCognitoAuthRefreshTokenAuth(t *testing.T) {
	accessToken := "access-token"
	idToken := "id-token"
	refreshToken := "refresh-token"
	testCases := []struct {
		Response cognitoidp.InitiateAuthOutput
		Expected map[string]string
	} {
		{
			Response: cognitoidp.InitiateAuthOutput {
				AuthenticationResult: &cognitoidp.AuthenticationResultType {
					AccessToken: aws.String(accessToken),
					ExpiresIn: aws.Int64(3600),
					IdToken: aws.String(idToken),
					RefreshToken: aws.String(refreshToken),
					TokenType: aws.String("Access"),
				},
			},
			Expected: map[string]string {
				"accessToken": accessToken,
				"idToken": idToken,
				"refreshToken": refreshToken,
			},
		},
	}

	for _, c := range testCases {
		cognitoAuth := CognitoAuth {
			Client: mockRefreshTokenAuth{Response: c.Response},
			UserPoolId: "mock-user-pool-id",
			AppClientId: "mock-app-client-id",
		}
		response, err := cognitoAuth.RefreshTokenAuth(refreshToken)
		if err != nil {
			t.Error("Error while signing in using refresh token.")
		}
		equal := reflect.DeepEqual(response, c.Expected)
		if ! equal {
			t.Error("Signin using refresh token failed.")
		}
	}
}


// >>> Test `CognitoAuth.ForgetPassword` method
type mockForgotPassword struct {
	cognitoidpiface.CognitoIdentityProviderAPI
	Response cognitoidp.ForgotPasswordOutput
}


func (mock mockForgotPassword) ForgotPassword(input *cognitoidp.ForgotPasswordInput) (*cognitoidp.ForgotPasswordOutput, error) {
	return &mock.Response, nil
}


func TestCognitoAuthForgotPassword(t *testing.T) {
	email := "user@email.com"
	testCases := []struct {
		Response cognitoidp.ForgotPasswordOutput
		Expected string
	} {
		{
			Response: cognitoidp.ForgotPasswordOutput {},
			Expected: fmt.Sprintf(
				"A verification code will be sent to email %s if a user exists with this email.", 
				email,
			),
		},
	}

	for _, c := range testCases {
		cognitoAuth := CognitoAuth {
			Client: mockForgotPassword {Response: c.Response},
			UserPoolId: "mock-user-pool-id",
			AppClientId: "mock-app-client-id",
		}
		response, err := cognitoAuth.ForgotPassword(email)
		if err != nil {
			t.Error("Error while getting verification code to reset password.")
		}
		if response != c.Expected {
			t.Error("Failed to get verification code to reset password.")
		}
	}
}
