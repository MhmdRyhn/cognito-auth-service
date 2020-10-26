package auth


import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	cognitoidp "github.com/aws/aws-sdk-go/service/cognitoidentityprovider" 
)


// Change password staying signed in
func ChangePassword(currentPassword string, newPassword string, accessToken string) (string, error) {
	changePasswordInput := &cognitoidp.ChangePasswordInput {
		PreviousPassword: aws.String(currentPassword),
		ProposedPassword: aws.String(newPassword),
		AccessToken: aws.String(accessToken),
    }

	_, err := client.ChangePassword(changePasswordInput)
	if err != nil {
		return "", err
	} else {
		return fmt.Sprintf("Password changed successfully."), err
	}
}
