package IdentityServer4_AccessTokenValidation

import (
	"github.com/getfloret/IdentityServer4.AccessTokenValidation/IdentityModel/oidc"
	"github.com/getfloret/IdentityServer4.AccessTokenValidation/options"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"strings"
)

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
				//todo
				errParse:= ParseJWT(token,oidc.DefaultKeyLoader)
				if(errParse == nil){
					c.Next()
				} else {
					log.Error(errParse)
					ErrInvalidToken.Write(c.Writer)
				}
			} else if options.GlobalAuthenticationOptions.SupportsIntrospection(){
				log.Trace("Token is a reference token and is supported.")
				c.Set(EffectiveSchemeKey,IntrospectionScheme)

				//todo
				//		return await Context.AuthenticateAsync(introspectionScheme);
				c.Next()
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