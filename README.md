# Cognito Auth Service
Auth service built using **Go** in **AWS Cognito**. To use this repository as your auth service, your **Cognito User Pool** must follow [**this**](https://github.com/MhmdRyhn/aws-cognito-iac) configurations.


## Available Functionalities
- User self signup
- Confirm signup
- Admin create user
- Force change password when user is created by admin
- Signin
- Get new Id-Token and Access-Token uaing Refresh-Token
- Reset password
- Change password staying signed in
- Admin delete user
- Create user group
- Admin add user to group
- Delete user group
- Admin remove user from group


## Exposed APIs
- Sign Up - `/auth/signup`
- Confirm Sign Up - `/auth/confirm-signup`
- Sign In - `/auth/signin`
- Refresh Token Auth - `/auth/refresh-token-auth`
- Forget Password - `/auth/forget-password`
- Confirm Forget Password - `/auth/confirm-forget-password`


## Run The Server
``` shell script
AWS_PROFILE={profile} AWS_REGION={region} USER_POOL={user-pool-id} APP_CLIENT_ID={app-client-id} go run cmd/main.go
```
