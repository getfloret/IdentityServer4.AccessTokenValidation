package IdentityServer4_AccessTokenValidation

import (
	"github.com/getfloret/IdentityServer4.AccessTokenValidation/IdentityModel"
	"github.com/getfloret/IdentityServer4.AccessTokenValidation/options"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/url"
)

const(
	AuthenticationScheme = "Bearer"
	AccessTokenParameter = "access_token"
	TokenItemsKey = "idsrv4:tokenvalidation:token"
	IdentityKey = "idsrv4:tokenvalidation:identity"
	EffectiveSchemeKey = "idsrv4:tokenvalidation:effective:"
	JWTScheme = "JWT"
	IntrospectionScheme = "Reference"
)



func IdentityServerAuthentication(configOptions func(*options.IdentityServerAuthenticationOptions)()) gin.HandlerFunc {
	configLog()
	configOptions(options.GlobalAuthenticationOptions)
	setupHandler()
	return HandleAuthenticate()
}

func configLog(){
	if(gin.Mode() == gin.DebugMode){
		log.SetLevel(log.TraceLevel)
	}else {
		log.SetLevel(log.WarnLevel)
		log.SetFormatter(&log.JSONFormatter{})
	}
}

func setupHandler(){
	oidc_discoverydoc_url, err := url.Parse(options.GlobalAuthenticationOptions.Authority + "/.well-known/openid-configuration")
	if(err != nil){
		log.Panic("Configuration Authority is not valid")
	}
	IdentityModel.DefaultKeyLoader = IdentityModel.NewCachingOpenIDProviderLoader(oidc_discoverydoc_url)
}