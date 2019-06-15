package IdentityModel

import (
	"github.com/getfloret/IdentityServer4.AccessTokenValidation/options"
)

// A KeyLoader fetches cryptographic keys and is able to lookup them up by ID or return the entire
// map of known keys
type KeyLoader interface {
	LoadKey(id string) (interface{}, error)
	Keys() map[string]interface{}
	OIDCConf() *options.OpenIdConnectConfiguration
}

var DefaultKeyLoader KeyLoader