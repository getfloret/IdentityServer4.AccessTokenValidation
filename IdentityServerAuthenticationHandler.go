package IdentityServer4_AccessTokenValidation

import (
	"github.com/getfloret/IdentityServer4.AccessTokenValidation/IdentityModel"
	"github.com/getfloret/IdentityServer4.AccessTokenValidation/options"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"strings"
)

var referenceTokenHandler = New(1000,options.GlobalAuthenticationOptions.CacheDuration)

func HandleAuthenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Trace("HandleAuthenticate called");
		token := AccessTokenFromRequest(c)
		removeToken := false
		defer func() {
			if(removeToken){
				delete(c.Keys, TokenItemsKey)
			}
		}()

		if(token != ""){
			log.Trace("Token found: {token}",token)
			removeToken = true
			c.Set(TokenItemsKey,token)

			// seems to be a JWT
			if(isJWTToken(token) && options.GlobalAuthenticationOptions.SupportsJwt()) {
				log.Trace("Token is a JWT and is supported.")
				c.Set(EffectiveSchemeKey,JWTScheme)
				mapClaims, errParse:= ParseJWT(token, IdentityModel.DefaultKeyLoader)
				if(errParse == nil){
					claims,tokenInfoError := IdentityModel.NewTokenInfo(mapClaims)
					if(tokenInfoError !=nil){
						panic(tokenInfoError)
					}
					log.Info(claims)
					c.Set(IdentityKey,claims)
					c.Next()
				} else {
					log.Error(errParse)
					ErrInvalidToken.Write(c.Writer)
					c.Abort()
				}
			} else if options.GlobalAuthenticationOptions.SupportsIntrospection(){
				log.Trace("Token is a reference token and is supported.")
				c.Set(EffectiveSchemeKey,IntrospectionScheme)
				mapClaims, errParse:= referenceTokenHandler.ParseReference(token)
				if(errParse == nil){
					claims,tokenInfoError := IdentityModel.NewTokenInfo(mapClaims)
					if(tokenInfoError !=nil){
						panic(tokenInfoError)
					}
					log.Info(claims)
					c.Set(IdentityKey,claims)
					c.Next()
				} else {
					log.Error(errParse)
					ErrInvalidToken.Write(c.Writer)
					c.Abort()
				}
			} else {
				log.Warn("Neither JWT nor reference tokens seem to be correctly configured for incoming token.")
				ErrInvalidToken.Write(c.Writer)
			}
		} else {
			ErrInvalidRequest.Write(c.Writer)
		}
	}
}

// AccessTokenFromRequest can be used to extract an Access Token from an http.Request
// via the standard headers/parameters
//  Ref:
//      https://tools.ietf.org/html/rfc6749#section-5.1
//      https://tools.ietf.org/html/rfc6750#section-2.1
func AccessTokenFromRequest(c *gin.Context) string {
	if h := c.Request.Header.Get("Authorization"); h != "" {
		if strings.HasPrefix(strings.ToLower(h), "bearer ") {
			return h[7:]
		}
	}

	return c.Request.FormValue(AccessTokenParameter)
}

func isJWTToken(token string) bool {
	if token == "" {
		return false
	}

	parts := strings.Split(token, ".")

	return len(parts) == 3
}