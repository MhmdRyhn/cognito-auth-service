package auth


import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	cognitoidp "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)


// Create a user group
func CreateUserGroup(groupName string, description string, precedence int64) (string, error) {
	userPoolId := os.Getenv("USER_POOL")
	createGroupInput := &cognitoidp.CreateGroupInput {
		UserPoolId: aws.String(userPoolId),
		GroupName: aws.String(groupName),
		Description: aws.String(description),
		Precedence: aws.Int64(precedence),
	}

	_, err := client.CreateGroup(createGroupInput)
	if err != nil {
		return "", err
	} else {
		return fmt.Sprintf("Group %s has been created successfully.", groupName), err
	}
}


// Add a user to a group
func AdminAddUserToGroup(username string, groupName string) (string, error) {
	userPoolId := os.Getenv("USER_POOL")
	adminAddUserToGroupInput := &cognitoidp.AdminAddUserToGroupInput {
		UserPoolId: aws.String(userPoolId),
		Username: aws.String(username),
		GroupName: aws.String(groupName),
	}

	_, err := client.AdminAddUserToGroup(adminAddUserToGroupInput)
	if err != nil {
		return "", err
	} else {
		return fmt.Sprintf(
			"User %s has been added to group %s successfully.", username, groupName,
		), err
	}
}


// Delete a user group
func DeleteUserGroup(groupName string) (string, error) {
	userPoolId := os.Getenv("USER_POOL")
	deleteGroupInput := &cognitoidp.DeleteGroupInput {
		UserPoolId: aws.String(userPoolId),
		GroupName: aws.String(groupName),
	}

	_, err := client.DeleteGroup(deleteGroupInput)
	if err != nil {
		return "", err
	} else {
		return fmt.Sprintf("Group %s has been deleted successfully.", groupName), err
	}
}


// Remove a user from a group
func AdminRemoveUserFromGroup(username string, groupName string) (string, error) {
	userPoolId := os.Getenv("USER_POOL")
	adminRemoveUserFromGroupInput := &cognitoidp.AdminRemoveUserFromGroupInput {
		UserPoolId: aws.String(userPoolId),
		Username: aws.String(username),
		GroupName: aws.String(groupName),
	}

	_, err := client.AdminRemoveUserFromGroup(adminRemoveUserFromGroupInput)
	if err != nil {
		return "", err
	} else {
		return fmt.Sprintf(
			"User %s has been removed from group %s successfully.", username, groupName,
		), err
	}
}
