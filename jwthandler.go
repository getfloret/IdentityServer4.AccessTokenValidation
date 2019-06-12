package IdentityServer4_AccessTokenValidation

import (
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/getfloret/IdentityServer4.AccessTokenValidation/IdentityModel/oidc"
)

var (
	// ErrMissingKeyID should be used when the kid attribute is missing in the JWT header
	ErrMissingKeyID = errors.New("Missing key Id in the JWT header")
	// ErrInvalidKeyID should be used when the content of the kid attribute is invalid
	ErrInvalidKeyID = errors.New("Invalid key Id in the JWT header")
)

func ParseJWT(tokenStr string, kl oidc.KeyLoader)(err error){
	token, err := jwt.Parse(tokenStr, getSignKey(kl))

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
		return nil
	} else {
		fmt.Println(err)
	}
	return err
}

func getSignKey(kl oidc.KeyLoader) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		switch token.Method.(type) {
		case *jwt.SigningMethodRSA, *jwt.SigningMethodECDSA:
			return loadKey(kl, token)
		default:
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
	}
}

func loadKey(kl oidc.KeyLoader, t *jwt.Token) (interface{}, error) {
	kid, has := t.Header["kid"]

	if !has {
		//todo review
		//return nil, ErrMissingKeyID
		kid = ""
	}

	id, ok := kid.(string)
	if !ok {
		return nil, ErrInvalidKeyID
	}

	return kl.LoadKey(id)
}
