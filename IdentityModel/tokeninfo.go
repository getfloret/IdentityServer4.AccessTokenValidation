package IdentityModel

import (
	"errors"
	"github.com/getfloret/IdentityServer4.AccessTokenValidation/consts/jwtclaimtypes"
	"github.com/getfloret/IdentityServer4.AccessTokenValidation/options"
	log "github.com/sirupsen/logrus"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type TokenInfo struct {
	Claims []JWTClaim
}

var (
	// ErrInvalidClaimScope should be used whenever the scope claim is invalid or missing in the JWT
	ErrInvalidClaimScope = errors.New("Invalid claim: scope")
	// ErrInvalidClaimRealm should be used whenever the scope realm is invalid or missing in the JWT
	ErrInvalidClaimRealm = errors.New("Invalid claim: realm")
	// ErrInvalidClaimSub should be used whenever the claim sub is invalid or missing in the JWT
	ErrInvalidClaimSub = errors.New("Invalid claim: sub")
	// ErrInvalidClaimAzp should be used whenever the claim azp is invalid or missing in the JWT
	ErrInvalidClaimAzp = errors.New("Invalid claim: azp")
	// ErrInvalidClaimExp should be used whenever the claim exp is invalid or missing in the JWT
	ErrInvalidClaimExp = errors.New("Invalid claim: exp")
)

func NewTokenInfo(claims jwt.MapClaims) (*TokenInfo, error) {
	if options.GlobalAuthenticationOptions.ClaimsProcessor != nil {
		transformClaims, transformErr := options.GlobalAuthenticationOptions.ClaimsProcessor.Process(claims)
		if transformErr == nil {
			claims = transformClaims
		} else {
			log.Error("Err when try to do custom transform token claims")
			return nil, transformErr
		}
	}

	//todo
	//claims.VerifyIssuer()
	return defaultNewTokenInfo(claims)
}

func defaultNewTokenInfo(claims jwt.MapClaims) (*TokenInfo, error) {
	jwtClaims := []JWTClaim{}
	for k, v := range claims {
		// due to a bug in identityserver - we need to be able to deal with the scope list both in array as well as space-separated list format
		if k == jwtclaimtypes.Scope {
			switch v.(type) {
			// it's an array
			case []interface{}:
				{
					for _, scopeV := range v.([]interface{}) {
						jwtClaims = append(jwtClaims, JWTClaim{ClaimType: jwtclaimtypes.Scope, ClaimValue: scopeV})
					}
				}
			// it's a string
			case string:
				{
					for item := range strings.Split(v.(string), " ") {
						jwtClaims = append(jwtClaims, JWTClaim{ClaimType: jwtclaimtypes.Scope, ClaimValue: item})
					}
				}
			}
		} else {
			jwtClaims = append(jwtClaims, JWTClaim{ClaimType: k, ClaimValue: v})
		}
	}

	return &TokenInfo{
		Claims: jwtClaims,
	}, nil
}
