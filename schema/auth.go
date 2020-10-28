package schema


type SignUpSchema struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=32,uppercase,lowercase,digit,punctuation"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,min=6,max=32,eqfield=Password,uppercase,lowercase,digit,punctuation"`
}


type ConfirmSignUpSchema struct {
	Email string `json:"email" validate:"required,email"`
	ConfirmationCode string `json:"confirmationCode" validate:"required,min=1"`
}
