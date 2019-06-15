package IdentityServer4_AccessTokenValidation

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/getfloret/IdentityServer4.AccessTokenValidation/IdentityModel"
	"github.com/getfloret/IdentityServer4.AccessTokenValidation/consts/jwtclaimtypes"
	"github.com/karlseguin/ccache"
	"time"
)

type referenceHandler struct {
	cache *ccache.Cache //LRU Cache
	cacheTTL time.Duration
}

var Err_TokenNotActive = errors.New("Token is not Active")

func New(cacheMaxSize int64,cacheTTL time.Duration) *referenceHandler{
	cache:= ccache.New(ccache.Configure().MaxSize(cacheMaxSize))
	return &referenceHandler{cache:cache,cacheTTL:cacheTTL}
}

func (h *referenceHandler) ParseReference(token string) (jwt.MapClaims, error){
	item := h.cache.Get(token)//todo user hash
	if item != nil {
		if !item.Expired() {
			return item.Value().(jwt.MapClaims), nil
		}
	}

	tokenResult, err := IdentityModel.Post(IdentityModel.DefaultKeyLoader.OIDCConf().IntrospectionEndpoint,&IdentityModel.TokenIntrospectionRequest{Token: token})
	if err==nil {
		if tokenResult.Active {
			if(h.cacheTTL>0){
				expiresIn := time.Unix(int64(tokenResult.Claims[jwtclaimtypes.Expiration].(float64)), 0).Sub(time.Now())
				var ttl time.Duration
				if expiresIn>h.cacheTTL{
					ttl = h.cacheTTL
				} else {
					ttl = expiresIn
				}
				h.cache.Set(token, tokenResult.Claims , ttl)
			}
			return tokenResult.Claims, nil
		} else {
			return nil,Err_TokenNotActive
		}
	 } else {
	 	return nil, err
	}
}