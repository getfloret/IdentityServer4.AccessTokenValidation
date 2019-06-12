package oidc

import "github.com/getfloret/IdentityServer4.AccessTokenValidation/IdentityModel"

// A KeyLoader fetches cryptographic keys and is able to lookup them up by ID or return the entire
// map of known keys
type KeyLoader interface {
	LoadKey(id string) (interface{}, error)
	Keys() map[string]interface{}
	OIDCConf() *IdentityModel.OpenIdConnectConfiguration
}

var DefaultKeyLoader KeyLoader