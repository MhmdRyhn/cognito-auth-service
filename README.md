# Cognito Auth Service
Auth service built using **Go** in **AWS Cognito**. To use this repository as your auth service, your **Cognito User Pool** must follow [**this**](https://github.com/MhmdRyhn/aws-cognito-iac) configurations.


## Run The Server
``` shell script
AWS_PROFILE={profile} AWS_REGION={region} USER_POOL={user-pool-id} APP_CLIENT_ID={app-client-id} go run cmd/auth/main.go
```
