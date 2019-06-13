package IdentityModel

import (
	"errors"
	"github.com/getfloret/IdentityServer4.AccessTokenValidation/options"
	log "github.com/sirupsen/logrus"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/getfloret/IdentityServer4.AccessTokenValidation/consts/jwtclaimtypes"
)

type TokenInfo struct {
	Claims map[string]string
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

// do user custom token transform & valid issuer .etc.
//https://github.com/IdentityModel/IdentityModel/blob/d95fd0713b4d2d93ee3a81c78ac970a76421294b/src/Client/Messages/TokenIntrospectionResponse.cs
//https://github.com/IdentityModel/IdentityModel.AspNetCore.OAuth2Introspection/blob/master/src/OAuth2IntrospectionHandler.cs

func NewTokenInfo(claims jwt.MapClaims) (*TokenInfo, error) {
	if options.GlobalAuthenticationOptions.ClaimsProcessor!=nil {
		transformClaims, transformErr := options.GlobalAuthenticationOptions.ClaimsProcessor.Process(claims)
		if(transformErr == nil){
			claims = transformClaims
		} else {
			log.Error("Err when try to do custom transform token claims")
			return nil,transformErr
		}
	}

	claims.VerifyIssuer()

	issuer, ok := ClaimAsString(t, jwtclaimtypes.Issuer)
	if ok {
		return
	}
	return defaultNewTokenInfo(t, timeBase)
}

// ALPS Proccess or Default TokenInfo
func defaultNewTokenInfo(t *jwt.Token, timeBase time.Time) (*processor.TokenInfo, error) {
	scopes, ok := ClaimAsStrings(t, jwtclaimtypes.Scope)
	if !ok {
		return nil, ErrInvalidClaimScope
	}

	sub, ok := ClaimAsString(t, jwtclaimtypes.Subject)
	if !ok {
		return nil, ErrInvalidClaimSub
	}

	realm, ok := ClaimAsString(t, JwtClaimRealm)
	if !ok {
		return nil, ErrInvalidClaimRealm
	}

	clientId := ""
	if claims, ok := t.Claims.(jwt.MapClaims); ok {
		_, has := claims[JwtClaimAzp]
		if has {
			clientId, ok = ClaimAsString(t, JwtClaimAzp)
			if !ok {
				return nil, ErrInvalidClaimAzp
			}
		}
	}

	exp, ok := ClaimAsInt64(t, jwtclaimtypes.Expiration)
	if !ok {
		return nil, ErrInvalidClaimExp
	}

	expiresIn := int(time.Unix(exp, 0).Sub(timeBase).Seconds())

	return &processor.TokenInfo{
		AccessToken: t.Raw,
		UID:         sub,
		GrantType:   "password",
		Scope:       scopes,
		Realm:       realm,
		ClientId:    clientId,
		TokenType:   "Bearer",
		ExpiresIn:   expiresIn,
	}, nil
}
