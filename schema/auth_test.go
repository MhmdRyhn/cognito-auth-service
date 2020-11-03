package schema

import (
	"reflect"
	"testing"

	"github.com/mhmdryhn/cognito-auth-service/schema/validation" // schemavalidation
)


// >>> Validate `schema.SignInSchema` schema
func TestValidateSignInSchema(t *testing.T) {
	var validSigninData, invalidSigninData []byte
	var errorDetails map[string]string
	var ok, equal bool

	// Test valid signin data
	validSigninData = []byte(`{"email": "user@mail.com", "password": "Password!23"}`)
	var signInSchema SignInSchema
	_, ok = schemavalidation.ValidateJSONData(validSigninData, &signInSchema)
	if !ok {
		t.Error("Incorrect Input.")
	}

	// Test invalid signin data-1
	invalidSigninData = []byte(`{"email": "usermail.com", "password": "Password!23"}`)
	errorDetails, ok = schemavalidation.ValidateJSONData(invalidSigninData, &signInSchema)
	if ok {
		t.Error("Input should be invalid.")
	}
	equal = reflect.DeepEqual(
		errorDetails, 
		map[string]string {"email": "Not a valid email address."},
	) 
	if !equal {
		t.Errorf(
			"Expected: %v | Found: %v",
			map[string]string {"email": "Not a valid email address."},
			errorDetails,
		)
	}

	// Test invalid signin data-2
	invalidSigninData = []byte(`{"email": "user@mail.com", "password": "password!23"}`)
	signInSchema = SignInSchema{}
	errorDetails, ok = schemavalidation.ValidateJSONData(invalidSigninData, &signInSchema)
	if ok {
		t.Error("Input should be invalid.")
	}
	equal = reflect.DeepEqual(
		errorDetails, 
		map[string]string {"password": "Must contain at least one capital letter."},
	) 
	if !equal {
		t.Errorf(
			"Expected: %v | Found: %v",
			map[string]string {"password": "Must contain at least one capital letter."},
			errorDetails,
		)
	}

	// Test invalid signin data-3
	invalidSigninData = []byte(`{"email": "user@mail.com", "password": "PASSWORD!23"}`)
	signInSchema = SignInSchema{}
	errorDetails, ok = schemavalidation.ValidateJSONData(invalidSigninData, &signInSchema)
	if ok {
		t.Error("Input should be invalid.")
	}
	equal = reflect.DeepEqual(
		errorDetails, 
		map[string]string {"password": "Must contain at least one small letter."},
	) 
	if !equal {
		t.Errorf(
			"Expected: %v | Found: %v",
			map[string]string {"password": "Must contain at least one small letter."},
			errorDetails,
		)
	}

	// Test invalid signin data-4
	invalidSigninData = []byte(`{"email": "user@mail.com", "password": "Password!@#"}`)
	signInSchema = SignInSchema{}
	errorDetails, ok = schemavalidation.ValidateJSONData(invalidSigninData, &signInSchema)
	if ok {
		t.Error("Input should be invalid.")
	}
	equal = reflect.DeepEqual(
		errorDetails, 
		map[string]string {"password": "Must contain at least one digit."},
	) 
	if !equal {
		t.Errorf(
			"Expected: %v | Found: %v",
			map[string]string {"password": "Must contain at least one digit."},
			errorDetails,
		)
	}

	// Test invalid signin data-5
	invalidSigninData = []byte(`{"email": "user@mail.com", "password": "Password123"}`)
	signInSchema = SignInSchema{}
	errorDetails, ok = schemavalidation.ValidateJSONData(invalidSigninData, &signInSchema)
	if ok {
		t.Error("Input should be invalid.")
	}
	equal = reflect.DeepEqual(
		errorDetails, 
		map[string]string {"password": "Must contain at least one punctuation."},
	) 
	if !equal {
		t.Errorf(
			"Expected: %v | Found: %v",
			map[string]string {"password": "Must contain at least one punctuation."},
			errorDetails,
		)
	}

	// Test invalid signin data-6
	invalidSigninData = []byte(`{"password": "Password!23"}`)
	signInSchema = SignInSchema{}
	errorDetails, ok = schemavalidation.ValidateJSONData(invalidSigninData, &signInSchema)
	if ok {
		t.Error("Input should be invalid.")
	}
	equal = reflect.DeepEqual(
		errorDetails, 
		map[string]string {"email": "Missing data for required field."},
	) 
	if !equal {
		t.Errorf(
			"Expected: %v | Found: %v",
			map[string]string {"email": "Missing data for required field."},
			errorDetails,
		)
	}

	// Test invalid signin data-7
	invalidSigninData = []byte(`{"email": "user@mail.com", "password": "Ab!1"}`)
	signInSchema = SignInSchema{}
	errorDetails, ok = schemavalidation.ValidateJSONData(invalidSigninData, &signInSchema)
	if ok {
		t.Error("Input should be invalid.")
	}
	equal = reflect.DeepEqual(
		errorDetails, 
		map[string]string {"password": "Length must be at least 6."},
	) 
	if !equal {
		t.Errorf(
			"Expected: %v | Found: %v",
			map[string]string {"password": "Length must be at least 6."},
			errorDetails,
		)
	}

	// Test invalid signin data-8
	invalidSigninData = []byte(`{"email": "user@mail.com", "password": "Password!23qwertyuiopasdfghjklzxcv"}`)
	signInSchema = SignInSchema{}
	errorDetails, ok = schemavalidation.ValidateJSONData(invalidSigninData, &signInSchema)
	if ok {
		t.Error("Input should be invalid.")
	}
	equal = reflect.DeepEqual(
		errorDetails, 
		map[string]string {"password": "Length can be at best 32."},
	) 
	if !equal {
		t.Errorf(
			"Expected: %v | Found: %v",
			map[string]string {"password": "Length can be at best 32."},
			errorDetails,
		)
	}
}
