package auth


import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	cognitoidp "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	cognitoidpiface "github.com/aws/aws-sdk-go/service/cognitoidentityprovider/cognitoidentityprovideriface"
)


// >>> Test `CognitoAuth.SignIn` method
type mockSignIn struct {
	cognitoidpiface.CognitoIdentityProviderAPI
	Response cognitoidp.InitiateAuthOutput
}


func (mock mockSignIn) InitiateAuth(input *cognitoidp.InitiateAuthInput) (*cognitoidp.InitiateAuthOutput, error) {
	return &mock.Response, nil
}


func TestCognitoAuthSignIn(t *testing.T) {
	testCases := []struct {
		Response cognitoidp.InitiateAuthOutput
		Expected map[string]string
	} {
		{
			Response: cognitoidp.InitiateAuthOutput {
				AuthenticationResult: &cognitoidp.AuthenticationResultType {
					AccessToken: aws.String("access-token"),
					ExpiresIn: aws.Int64(3600),
					IdToken: aws.String("id-token"),
					RefreshToken: aws.String("refresh-token"),
					TokenType: aws.String("Access"),
				},
			},
			Expected: map[string]string {
				"accessToken": "access-token",
				"idToken": "id-token",
				"refreshToken": "refresh-token",
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
