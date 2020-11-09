// Package for utilities 

package utils


import (
	"errors"
	"fmt"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/lestrrat-go/jwx/jwk"
)


var (
	awsRegion string = os.Getenv("AWS_REGION")
	userPoolId string = os.Getenv("USER_POOL_ID")
	appClientId string = os.Getenv("APP_CLIENT_ID")
	publicKeysURL string = fmt.Sprintf(
		"https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json", 
		awsRegion, userPoolId,
	)
	issuer string = fmt.Sprintf(
		"https://cognito-idp.%s.amazonaws.com/%s",
		awsRegion, userPoolId,
	)
)


type JWTValidator struct {
	// The claims found in the token right after the 
	// token is parsed successfully.
	claims jwt.MapClaims
}


// `KeyFunc` in `jwt.Parse`. It validates the signing method 
// (RS256 in this caase). Then, it build token key using JWKs.
func (self *JWTValidator) KeyFunction (token *jwt.Token) (interface{}, error) {
	// Validate signing method/algorithm
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
		return nil, errors.New("Failed to parse public key")
	}
	// Search `kid` in JWKs
	keys := publicKeySet.LookupKeyID(kid);
	if len(keys) == 0 {
			return nil, errors.New("Key not found.")
	}
	// Create token key
	var tokenKey interface{}
	if err := keys[0].Raw(&tokenKey); err != nil {
		return nil, errors.New("Failed to create token key.")
	}
	return tokenKey, nil
}


// Validate token, its signature, basic claims etc. 
// and return all the claims present in the token.
func (self *JWTValidator) ParseJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, self.KeyFunction)
	if token != nil && token.Valid {
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			self.claims = claims
			return claims, nil
		}
		return nil, err
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors & jwt.ValidationErrorMalformed != 0 {
			return nil, errors.New("This is not even a token. Please make sure you are using a valid token.")
		} else {
			return nil, err
		}
	}
	return nil, err
}


// Verify `client_id` in the token.
func (self *JWTValidator) VerifyClient(client string) bool {
	return self.claims["client_id"] == client
}


// Verifies the token. It verifies:
// - If it is actually a token
// - Signing method (RS256)
// - Expiration
// - Issuer
// - Client
func (self *JWTValidator) Validate(tokenString string) (jwt.MapClaims, error) {
	claims, err := self.ParseJWT(tokenString)
	if err != nil {
		return nil, err
	}
	if ok := claims.VerifyIssuer(issuer, true); !ok {
		return nil, errors.New("Unknown issuer.")
	}
	if ok := self.VerifyClient(appClientId); !ok {
		return nil, errors.New("Unknown client.")
	}
	return claims, err
}
