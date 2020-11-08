package utils


import (
	"fmt"
	"errors"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/lestrrat-go/jwx/jwk"
)


var publicKeysURL string = fmt.Sprintf(
	"https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json", 
	os.Getenv("AWS_REGION"),
	os.Getenv("USER_POOL_ID"),
	// "1234567",
)


// type JWTValidatorMethod interface {
// 	func KeyFunction (*jwt.Token) (interface{}, error)
// }


// type JWTValidator struct {
// 	tokenString string
// }


func KeyFunction (token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	// Get `kid` from token's header
	kid, ok := token.Header["kid"].(string)
	if !ok {
			return nil, errors.New("kid header not found")
	}
	// Get JWKs (JSON Web Keys)
	publicKeySet, err := jwk.Fetch(publicKeysURL)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to parse key: %s", err))
	}
	// Search `kid` in JWKs
	keys := publicKeySet.LookupKeyID(kid);
	if len(keys) == 0 {
			return nil, fmt.Errorf("key %v not found", kid)
	}

	var tokenKey interface{}
	if err := keys[0].Raw(&tokenKey); err != nil {
		return nil, errors.New("Failed to create token key")
	}

	return tokenKey, nil
}


func ParseJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, KeyFunction,)

	if token.Valid {
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			return claims, nil
		}
		return nil, err
	// } else if ve, ok := err.(*jwt.ValidationError); ok {
	// 	// fmt.Println(token)
	// 	// fmt.Println(ve)
	// 	if ve.Errors & jwt.ValidationErrorMalformed != 0 {
	// 		return nil, err
	// 	} else if ve.Errors & (jwt.ValidationErrorExpired | jwt.ValidationErrorNotValidYet) != 0 {
	// 		// Token is either expired or not active yet
	// 		return nil, err
	// 	} else {
	// 		return nil, err
	// 	}
	} else {
		return nil, err
	}
}
