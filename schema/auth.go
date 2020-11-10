package schema


type SignUpSchema struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=32,uppercase,lowercase,digit,punctuation"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,min=6,max=32,eqfield=Password,uppercase,lowercase,digit,punctuation"`
}


type ResendConfirmationCodeSchema struct {
	Email string `json:"email" validate:"required,email"`
}


type ConfirmSignUpSchema struct {
	Email string `json:"email" validate:"required,email"`
	ConfirmationCode string `json:"confirmationCode" validate:"required,min=1"`
}


type SignInSchema struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=32,uppercase,lowercase,digit,punctuation"`
}


type RefreshTokenAuthSchema struct {
	RefreshToken string `json:"refreshToken" validate:"required,min=1"`
}


type ForgetPasswordSchema struct {
	Email string `json:"email" validate:"required,email"`
}


type ConfirmForgetPasswordSchema struct {
	Email string `json:"email" validate:"required,email"`
	ConfirmationCode string `json:"confirmationCode" validate:"required,min=1"`
	Password string `json:"password" validate:"required,min=6,max=32,uppercase,lowercase,digit,punctuation"`
}


type ChangePasswordSchema struct {
	CurrentPassword string `json:"currentPassword" validate:"required,min=6,max=32,uppercase,lowercase,digit,punctuation"`
	NewPassword string `json:"newPassword" validate:"required,min=6,max=32,uppercase,lowercase,digit,punctuation"`
	AccessToken string `json:"accessToken" validate:"required,min=1"`
}
